package application

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/errors"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/logging"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/services"
	metrics "github.com/isteshkov/highload-social-network/third_party/metrics"
)

var errorProducer = errors.NewProducer("general error")

type API interface {
	Run(address ...string) error
}

func BuildApplication(l logging.Logger, s *services.Services, m metrics.Metrics, cfg Config) *Application {
	app := &Application{
		l:        l,
		services: s,
		cfg:      cfg,
		metrics:  m,
	}

	app.publicApi = app.buildPublicApi()
	app.metricApi = app.buildMetricApi()
	app.profilingApi = buildProfilingApi()

	return app
}

type Application struct {
	l            logging.Logger
	services     *services.Services
	metrics      metrics.Metrics
	publicApi    API
	profilingApi API
	metricApi    API
	cfg          Config
}

func (a *Application) SetLogger(l logging.Logger) {
	a.l = l
	a.services = a.services.WithLogger(l)
}

func (a *Application) ListenAndServe() {
	done := make(chan bool)
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	go func() {
		sig := <-signals
		a.l.WithField("signal", sig).Info("get interrupting signal")
		done <- true
		return
	}()

	if a.cfg.ProfilingApiPort != "" {
		go func() {
			err := a.profilingApi.Run(a.cfg.ProfilingApiPort)
			if err != nil {
				a.l.Error(errorProducer.Wrap(err))
			}
			done <- true
			return
		}()
		a.l.Info("Run profiling API on port %s", a.cfg.ProfilingApiPort)
	}

	if a.cfg.MetricApiPort != "" {
		go func() {
			err := a.metricApi.Run(a.cfg.MetricApiPort)
			if err != nil {
				a.l.Error(errorProducer.Wrap(err))
			}
			done <- true
			return
		}()
		a.l.Info("Run metric API on port %s", a.cfg.MetricApiPort)
	}

	go func() {
		err := a.publicApi.Run(a.cfg.PublicApiPort)
		if err != nil {
			a.l.Error(errorProducer.Wrap(err))
		}
		done <- true
		return
	}()

	a.l.WithField("pid", os.Getpid()).Info("Run public API on port %s", a.cfg.PublicApiPort)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range done {
			a.l.Info("Shutdown application")
			time.Sleep(time.Millisecond)
			return
		}
	}()

	wg.Wait()
}

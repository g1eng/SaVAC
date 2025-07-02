package cloud_actions

import (
	"context"
	"fmt"
	"strings"

	"github.com/g1eng/savac/cmd/helper"
	"github.com/g1eng/savac/pkg/cloud/model/apprun"
	"github.com/g1eng/savac/pkg/core"
	"github.com/urfave/cli/v3"
)

func buildAppRunRegistryCredentials(cmd *cli.Command) (*apprun.NullableString, *apprun.NullableString) {
	u, p := cmd.String("user"), cmd.String("password")
	up, pp := apprun.NewNullableString(&u), apprun.NewNullableString(&p)
	if u == "" {
		up = apprun.NewNullableString(nil)
	}
	if p == "" {
		pp = apprun.NewNullableString(nil)
	}
	return up, pp
}

func buildAppRunHealthcheckProbe(cmd *cli.Command) *apprun.PostApplicationBodyComponentProbe {
	prob := apprun.NewPostApplicationBodyComponentProbe()
	prob.HttpGet.Set(&apprun.PostApplicationBodyComponentProbeHttpGet{
		Path:    cmd.String("path"),
		Port:    cmd.Int32("health-port"),
		Headers: nil,
	})
	return prob
}

func (g *CloudActionGenerator) getAppRunApplicationWithName(ctx context.Context, name string) (apprun.HandlerListApplicationsData, error) {
	res, _, err := g.ApiClient.AppRunAPI.DefaultApi.ListApplications(ctx).Execute()
	if err != nil {
		return apprun.HandlerListApplicationsData{}, err
	}
	return core.MatchResourceWithName(res.Data, name)
}

// mapAppRunCmdToPostApplicationBody
func mapAppRunCmdToPostApplicationBody(cmd *cli.Command) apprun.PostApplicationBody {
	name := cmd.String("name")
	cpu := fmt.Sprintf("%0.1f", cmd.Float64("cpu"))
	memory := fmt.Sprintf("%dMi", cmd.Int32("memory"))
	port := cmd.Int32("port")
	i := cmd.Args().First()
	s := strings.Split(i, "/")[0]
	timeout := cmd.Int32("timeout")
	up, pp := buildAppRunRegistryCredentials(cmd)
	//var prob *apprun.PostApplicationBodyComponentProbe
	//prob = buildAppRunHealthcheckProbe(cmd)

	return apprun.PostApplicationBody{
		Name:           name,
		TimeoutSeconds: timeout,
		Port:           port,
		MinScale:       0,
		MaxScale:       2,
		Components: []apprun.PostApplicationBodyComponent{
			{
				Name:      name + "-app",
				MaxCpu:    cpu,
				MaxMemory: memory,
				DeploySource: apprun.PostApplicationBodyComponentDeploySource{
					ContainerRegistry: &apprun.PostApplicationBodyComponentDeploySourceContainerRegistry{
						Image:    i,
						Server:   *apprun.NewNullableString(&s),
						Username: *up,
						Password: *pp,
					},
				},
				Env: nil,
				//Probe: *apprun.NewNullablePostApplicationBodyComponentProbe(prob),
			},
		},
	}
}

func (g *CloudActionGenerator) GenerateAppRunApplicationGetAction(ctx context.Context, cmd *cli.Command) error {
	appId := cmd.Args().First()
	if cmd.IsSet("version") {
		versionId := cmd.String("version")
		// Because the API endpoint is unstable at now, we do not decode the response
		res, _, err := g.ApiClient.AppRunAPI.DefaultApi.GetApplicationVersion(ctx, appId, versionId).Execute()
		if err != nil {
			return err
		}
		return helper.PrintJson(res)
	}
	if app, err := g.getAppRunApplicationWithName(ctx, appId); err == nil {
		appId = app.Id
	}

	// Because the API endpoint is unstable at now, we do not decode the response
	res, _, err := g.ApiClient.AppRunAPI.DefaultApi.GetApplication(ctx, appId).Execute()
	if err != nil {
		return err
	}
	return helper.PrintJson(res)
}

func mapAppRunCmdToPatchApplicationBody(cmd *cli.Command) apprun.PatchApplicationBody {
	postAppBody := mapAppRunCmdToPostApplicationBody(cmd)
	var components []apprun.PatchApplicationBodyComponent
	for _, component := range postAppBody.Components {
		depSources := apprun.PatchApplicationBodyComponentDeploySource{
			ContainerRegistry: &apprun.PatchApplicationBodyComponentDeploySourceContainerRegistry{
				Server:   component.DeploySource.ContainerRegistry.Server,
				Image:    component.DeploySource.ContainerRegistry.Image,
				Username: component.DeploySource.ContainerRegistry.Username,
				Password: component.DeploySource.ContainerRegistry.Password,
			},
		}
		components = append(components, apprun.PatchApplicationBodyComponent{
			Name:         component.Name,
			MaxCpu:       component.MaxCpu,
			MaxMemory:    component.MaxMemory,
			DeploySource: depSources,
			Env:          nil,
			Probe:        apprun.NullablePatchApplicationBodyComponentProbe{},
		})
	}

	atav := true
	return apprun.PatchApplicationBody{
		TimeoutSeconds:      &postAppBody.TimeoutSeconds,
		Port:                &postAppBody.Port,
		MinScale:            &postAppBody.MinScale,
		MaxScale:            &postAppBody.MaxScale,
		Components:          components,
		AllTrafficAvailable: &atav,
	}
}

func (g *CloudActionGenerator) GenerateAppRunApplicationListAction(ctx context.Context, cmd *cli.Command) error {
	if cmd.IsSet("versions") {
		appId := cmd.String("versions")
		req, _, err := g.ApiClient.AppRunAPI.DefaultApi.ListApplicationVersions(ctx, appId).Execute()
		if err != nil {
			return err
		}
		return helper.PrintJson(req)
	}
	res, _, err := g.ApiClient.AppRunAPI.DefaultApi.ListApplications(ctx).Execute()
	if err != nil {
		return err
	}
	return helper.PrintJson(res)
}

func (g *CloudActionGenerator) GenerateAppRunApplicationCreateAction(ctx context.Context, cmd *cli.Command) error {
	req := g.ApiClient.AppRunAPI.DefaultApi.PostApplication(ctx)
	req = req.PostApplicationBody(mapAppRunCmdToPostApplicationBody(cmd))

	res, _, err := req.Execute()
	if err != nil {
		return err
	}
	return helper.PrintJson(res)
}

func (g *CloudActionGenerator) GenerateAppRunApplicationCreateVersionAction(ctx context.Context, cmd *cli.Command) error {
	appId := cmd.String("app-id")
	if app, err := g.getAppRunApplicationWithName(ctx, appId); err == nil {
		appId = app.Id
	}
	req := g.ApiClient.AppRunAPI.DefaultApi.PatchApplication(ctx, appId)
	req = req.PatchApplicationBody(mapAppRunCmdToPatchApplicationBody(cmd))
	_, _, err := req.Execute()
	if err != nil {
		return err
	}

	return nil
}

func (g *CloudActionGenerator) GenerateAppRunApplicationDeleteAction(ctx context.Context, cmd *cli.Command) error {
	appId := cmd.Args().First()
	if cmd.IsSet("version") {
		versionId := cmd.String("version")
		// Because the API endpoint is unstable at now, we do not decode the response
		_, err := g.ApiClient.AppRunAPI.DefaultApi.DeleteApplicationVersion(ctx, appId, versionId).Execute()
		if err != nil {
			return err
		}
		return nil
	}
	if app, err := g.getAppRunApplicationWithName(ctx, appId); err == nil {
		appId = app.Id
	}

	// Because the API endpoint is unstable at now, we do not decode the response
	_, err := g.ApiClient.AppRunAPI.DefaultApi.DeleteApplication(ctx, appId).Execute()
	if err != nil {
		return err
	}
	return nil
}

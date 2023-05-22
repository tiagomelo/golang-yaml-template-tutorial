# golang-yaml-template-tutorial

Go templating is a powerful feature provided by the Go programming language that allows you to generate text output by replacing placeholders (variables) in a template with their corresponding values. It's a convenient way to generate dynamic content, including YAML files.

[Helm](https://helm.sh/), a package manager for [Kubernetes](https://kubernetes.io/), utilizes Go templating to generate [Kubernetes manifest files](https://kubernetes.io/docs/concepts/cluster-administration/manage-deployment/) from templates. 

In this tutorial we'll see how we can use Go templating to replace values in a YAML file with values from another YAML file, similar to what Helm does.

## running it

```
$ make run
```

It will generate `parsed/parsed.yaml` file as a result of replacing placeholders in `template/template.yaml` with values coming from `template/values.yaml`.

## unit tests

```
$ make test
```

## unit test coverage

```
$ make coverage
```
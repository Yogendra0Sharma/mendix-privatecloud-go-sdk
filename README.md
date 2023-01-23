![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Yogendra0Sharma/mendix-privatecloud-go-sdk?filename=go.mod)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)


# Mendix Private Cloud Go SDK Version 0.0.1
Official Docs: <https://docs.mendix.com/apidocs-mxsdk/apidocs/private-cloud-deploy-api/>
# Using the SDK
## 1 Authentication
Authentication for the API uses a Personal Access Token (PAT).

## 1.1 Generating a PAT
Go to <https://warden.mendix.com/> and follow the instructions in Create a [Personal Access Token](https://docs.mendix.com/developerportal/community-tools/warden/
) with Warden. Select the following as scopes:

*set MX_TOKEN=<GENERATED_PAT>*
*mx:deployment:read* -- to perform GET operations

*mx:deployment:write -- to perform all operations (GET, POST, PUT, and DELETE)
Store the generated value {GENERATED_PAT} somewhere safe so you can use it to authorize your Mendix for Private Cloud API calls.

## 1.2 Using the PAT 
Set Environment variable  MX_TOKEN

```
set MX_TOKEN=<GENERATED_PAT>
```

## 2 Run the tests

```
make test

```

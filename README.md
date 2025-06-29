# benchmark

## requirements

- Docker
- Make

## getting started

### codegen

The following command generates the SDKs for the API 
endpoints defined in the ProtoBuf files available in `_schema/protos/`.

It creates an SDK for each technology stack under the `api/*-sdk` directories.

```bash
make api-codegen
```
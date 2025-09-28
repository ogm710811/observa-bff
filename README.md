# Observa Tech Health Dashboard — BFF Aggregator

## Overview

This project implements the **Backend-for-Frontend (BFF) Aggregator** for the Observa Tech Health Dashboard. It serves as a single API boundary for the dashboard SPA (single-page application), orchestrating and normalizing data from both external (Observe, CI/CD, PagerDuty) and internal sources (service catalog, deploy streams, configuration). Built for AWS Lambda using **Go 1.25.x** and the AWS Lambda Go SDK, it provides a scalable, observable, and secure API layer.

For an executive architecture diagram, open `BFF-AGGREGATOR-ARCHITECTURE.drawio` with diagrams.net.

---

## Tech Stack

- **Language**: Go >=1.25.x
- **Runtime**: AWS Lambda (Go) via AppSync (optionally AWS API Gateway)
- **Infrastructure**: AWS CloudFormation (see [`observa-bff-lambda.yaml`](infra/observa-bff-lambda.yaml)), S3, IAM, VPC, ElastiCache (Redis), CloudWatch, AppConfig, Secrets Manager
- **CI/CD**: GitHub Actions (see [`deploy.yml`](.github/workflows/deploy.yml))
- **Observability**: OpenTelemetry traces, RED metrics, structured logs

---

## Architecture & Responsibilities

- **Single GraphQL API endpoint:** Aggregates data for dashboard queries (metrics, deploys, teams, incidents, etc). Also supports paginated and view-aware payloads.
- **Data orchestration:** Fans out queries in-parallel to data sources: Observe, service catalog, deploy feeds, CI/CD, configuration, identity providers.
- **Caching:** In-memory per-Lambda, cross-instance Redis, and query batching, all with freshness and watermarks.
- **Access control:** OIDC authentication, JWT claims, and row-level team filtering.
- **Configuration management:** Thresholds and visual settings in AppConfig; tokens in Secrets Manager.
- **Resilience:** Partial responses with status envelopes, circuit breakers, hedged queries, and cache fallback.

---

## Deployment

Deployments are managed via GitHub Actions and CloudFormation. Ephemeral/PR environments are supported via environment suffixing.

**Key files:**
- [`deploy.yml`](.github/workflows/deploy.yml): Build, test, and deploy workflow (including linting CloudFormation).
- [`observa-bff-lambda.yaml`](infra/observa-bff-lambda.yaml): CloudFormation template for the Lambda function, IAM role, etc.

**To deploy:**
1. Open a Pull Request.
2. GitHub Actions builds, uploads a ZIP to S3, and deploys/updates the CloudFormation stack.
3. PR deployments auto-suffix resources to avoid conflict.

---

## Local Development

- **Entry point:** `main.go` (lambda handler)
- **Run locally:** You can build and test code using Go tooling. For Lambda-like local invocations, use [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/) or [AWS Lambda Runtime Interface Emulator](https://github.com/aws/aws-lambda-runtime-interface-emulator).

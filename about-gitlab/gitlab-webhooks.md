# Gitlab Webhooks Documentation

https://docs.gitlab.com/ee/user/project/integrations/webhooks.html

# General Webhook Notes

* Webhooks are a custom HTTP callback that we define.

> Webhooks are custom HTTP callbacks that you define. They are usually triggered by an event, such as pushing code to a repository or posting a comment on an issue. When the event occurs, the source app makes an HTTP request to the URI configured for the webhook. The action to take may be anything. For example, you can use webhooks to:

> Trigger continuous integration (CI) jobs, update external issue trackers, update a backup mirror, or deploy to your production server

> Send a notification to Slack every time a job fails.

> Integrate with Twilio to be notified via SMS every time an issue is created for a specific project or group in GitLab.

> Automatically assign labels to merge requests.

## Configuring a Webhook

To configure a webhook for a project or group:

1. In your project or group, on the left sidebar, select Settings > Webhooks.
2. In URL, enter the URL of the webhook endpoint. The URL must be percent-encoded if it contains one or more special characters.
3. In Secret token, enter the secret token to validate payloads.
4. In the Trigger section, select the events to trigger the webhook.
5. Optional. Clear the Enable SSL verification checkbox to disable SSL verification.
6. Select Add webhook.

* There are other parts of the webhook which can be 

## Testing Webhooks

Webhooks can be [tested](https://docs.gitlab.com/ee/user/project/integrations/webhooks.html#test-a-webhook)

## Webhook Payload

* Webhooks can have a payload, e.g. in the form of JSON.

* Payloads can be [validated with a secret token](https://docs.gitlab.com/ee/user/project/integrations/webhooks.html#validate-payloads-by-using-a-secret-token).

## Event Types

* Webhooks can fire on different [event types](https://docs.gitlab.com/ee/user/project/integrations/webhook_events.html)

E.g. if an issue is created, or if a merge request event is created, merged, closed, etc.


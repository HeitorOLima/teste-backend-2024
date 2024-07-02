# frozen_string_literal: true

# Application consumer from which all Karafka consumers should inherit
# You can rename it if it would conflict with your current code base (in case you're integrating
# Karafka with other frameworks)
class ApplicationConsumer < Karafka::BaseConsumer
  include Karafka::Backends::Retry
  self.retry_policy = { interval: 5, max_retries: 10 }
end

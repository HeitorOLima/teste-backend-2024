# frozen_string_literal: true
class GoConsumer < ApplicationConsumer
  def consume
    messages.each do |message| 
      payload = message.payload.with_indifferent_access
      upsert_product(payload)
    end
  end

  private

  def upsert_product(payload)
    product_params = payload.merge(is_api: false)
    ::Services::Api::V1::Products::Upsert.new(product_params, nil).execute
  end 
end
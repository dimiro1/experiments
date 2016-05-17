require 'bundler/setup'
require 'hanami/router'

class Index
  def call(env)
    [200, { 'Content-Type' => 'text/plain' }, ['Index'] ]
  end
end

class Hello
  def call(env)
    [200, { 'Content-Type' => 'text/plain' }, ["Hello #{env['router.params'][:name]}"] ]
  end
end

app = Hanami::Router.new do
  get '/',            to: Index.new
  get '/hello/:name', to: Hello.new
end

Rack::Server.start app: app
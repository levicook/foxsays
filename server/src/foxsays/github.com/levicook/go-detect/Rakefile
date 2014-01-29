
namespace :go do

  desc "Run go install"
  task install: %w(go:test) do
    system('go install')
    exit(1) unless $?.success?
  end

  desc "Run go test ./..."
  task :test do
		puts "-"*10
    system('go test ./...')
    exit(1) unless $?.success?
  end

  desc "Run go fmt ./..."
  task :fmt do
    system('go fmt ./...')
    exit(1) unless $?.success?
  end

end

desc 'Test all the things'
task test: %w(go:test)

task default: %w(install)

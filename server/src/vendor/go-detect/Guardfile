interactor :off

guard 'shell' do
	watch(%r{\.go$}) { |m| system('rake go:install') }
end

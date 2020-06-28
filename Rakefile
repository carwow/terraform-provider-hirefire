# frozen_string_literal: true

namespace :build do
  task :release, [:version, :os, :arch] do |_t, args|
    version = args[:version]
    os = args[:os]
    arch = args[:arch]

    if version.nil? || os.nil? || arch.nil?
      abort "Requires version, os, and arch argument.\n" \
            'For example: rake build:release[0.1.0,linux,amd64]'
    end

    sh('env', "GOOS=#{os}", "GOARCH=#{arch}", 'go', 'build', '-o',
       "terraform-provider-hirefire_v#{version}_#{os}_#{arch}")
  end
end

desc 'Builds binaries for release (and creates GitHub release if hub is available)'
task :release, [:version] do |_t, args|
  version = args[:version]

  if version.nil?
    abort "Requires version argument.\n" \
          'For example: rake release[0.1.0]'
  end

  [
    %w[linux amd64]
  ].each do |os_and_arch|
    Rake::Task['build:release'].invoke(version, *os_and_arch)
  end

  hub_is_available = !`which hub`.empty?

  sh 'hub', 'release', 'create', "v#{version}" if hub_is_available

  puts
  puts '#' + '-' * 19 + '#'
  puts '# Still left to do: #'
  puts '#' + '-' * 19 + '#'
  puts
  unless hub_is_available
    puts " * Create a GitHub release:\n" \
         '   https://github.com/carwow/terraform-provider-hirefire/releases/new'
    puts
  end
  puts " * Upload binary files to GitHub release:\n" \
       "   https://github.com/carwow/terraform-provider-hirefire/releases/tag/v#{version}"
  puts
end

task default: :release

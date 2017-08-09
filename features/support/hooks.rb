require 'aruba'

def removeAllContainers()
    containers = `docker container ls --all --quiet`
    unless containers.to_s.strip.empty?
        `docker container rm -f #{containers}`
    end
end

Before do
    removeAllContainers
end

After do
    removeAllContainers
end

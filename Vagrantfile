# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.box_download_insecure = true
  config.vm.box = "debian/jessie64"
  config.vm.box_check_update = true
  config.vm.network "forwarded_port", guest: 6697, host: 6697

  # Create a private network, which allows host-only access to the machine
  # using a specific IP.
  # config.vm.network "private_network", ip: "192.168.33.10"

  config.vm.provider "virtualbox" do |vb|
    vb.gui = false
    vb.name = "tad_dev"
    vb.memory = "512"
  end
end

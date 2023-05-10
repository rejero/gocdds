# gocdds
Cyclone DDS Go Interface

Pulled from: https://github.com/ami-GS/go-cdds

The wrapper has been re-organised and syntax has been updated to the latest CycloneDDS version.


## Introduction
GoCDDS is a rough implementation of a Go wrapper for Cyclone DDS.

Golang is a statically typed modern programming language that has increasingly become more popular among modern day software engineers. Other than Golang's own Gobot (a robotics framework), modern industrial robotics tools such as ROS2 does not have Golang support.

Instead of implementing a Golang wrapper for ROS2, directly having a Golang interface with the DDS (middleware underlying ROS2) could prove to be a more efficient mode of data transmission between nodes without having an additional layer between each ROS2 package coded in Golang and ROS2 itself.

As Robotics Engineers are difficult to train up, it is important to lower the barrier of entry by creating tools that allows software engineers from another specialisation to cross over to Robotics.


## Installation
The installation has been tested on Ubuntu 22.04.1 (Jammy Jellyfish)

1. Install the iceoryx libraries either from source or from apt.
```
sudo apt install libiceoryx-binding-c1 libiceoryx-binding-c-dev libiceoryx-platform1 libiceoryx-posh1 libiceoryx-posh-config1 libiceoryx-posh-dev libiceoryx-posh-gateway1 libiceoryx-posh-roudi1 libiceoryx-utils1 libiceoryx-utils-dev
```
2. Install CycloneDDS either from source or from apt.
```
sudo apt install cyclonedds-dev cyclonedds-doc cyclonedds-tools libcycloneddsidl0 libddsc0
```
3. Clone the git repository into your Go workspace
```
git clone https://github.com/rejero/gocdds
```
4. Build the CycloneDDS examples with their Go versions
```
cd gocdds
pushd examples/helloworld/
cmake .
cmake --build .
popd
pushd examples/roundtrip/
cmake .
cmake --build .
popd
pushd examples/throughput/
cmake .
cmake --build .
popd
```

### Running examples as tests
Test the installation by running some few examples
#### Helloworld
1. Navigate to gocdds/examples/helloworld folder
```
pushd examples/helloworld/
```
2. Run the command below in terminal to start the Go subscriber
```
go run subscriber/subscribe.go
```

You should see the following output:

```
=== [Subscriber] Waiting for sample ...
```
3. Run the command below in another terminal to start the Go publisher
```
go run publisher/publish.go
```

You should see that the publisher has sent a sample and terminal outputs the following

```
=== [Publisher] Waiting for a reader to be discovered ...
=== [Publisher] Writing : 
Message (12343, {"Name":"cyclone", "Age":22})
```

Go back to the subscriber terminal and an additional line should be shown to show that the sample has been received.

```
=== [Subscriber] Received : Message 0:(12343, {"Name":"cyclone", "Age":22})
```


## Warnings
Currently several methods has issue when its object is deleted via participant.Delete().
The methods are mainly have loop in it.
for instance, Reader.BlockAllocRead() and Writer.SearchTopic()


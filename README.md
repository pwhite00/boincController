# boincController : An app that enables control of a headless BOINC docker container.

What is boincController ?

 - BoincController is a small app written in golang that enables easy setup , launch and control of boinc containers.
 - BoincController is the actual boinc docker container.
 - A configurtation template for controlling what boinc containers to run.

Ok but What is BOINC ??
  [BOINC](https://boinc.berkeley.edu/) is an open computing platform designed to use shared computing resources. It allows
  volunteers to use extra cou cycles to solve hard computuing problems by donting unused time on a local machine.
  
  Some of the projects that are avalable to volunteer on are:
    * [SETI@home](https://setiathome.berkeley.edu/)
    * [IBM World Community Grid](https://worldcommunitygrid.org/)
    * [Einstein@Home](https://einsteinathome.org/)
   
   
   
   How to Use boiuncController:
   
   1) Make sure you have docker installed on your machine.
   2) Get the container on your machine.
      a) Pull from dockerhub.com
      <pre>
      $ docker pull pwhite00/boinc_client:latest
      </pre>
      b) use the Dockerfile to locally build your container. (included in your forked copy of the repo.)
      <pre>
      $ git clone git@github.com:{you}/boincController.git
      $ cd boincController
      $ docker build -t boinc_client .
      $ docker images ls | grep boinc_client
      </pre>
   2) Edit your config file:
   Exmaple config:
   <pre>
   {
      "containers": [
        {
          "baseName": "worldComputeGrid",
          "runningCount": 2,
          "projectTarget": "http://www.worldcommunitygrid.org/",
          "projectKey": "QVk4boOESndk8dXuqkfM0v8WYv5fflhv"
        },
        {
          "baseName": "setiAtHome",
          "runningCount": 2,
          "projectTarget": "http://setiathome.berkeley.edu/",
          "projectKey": "0VGTCrsHC9f4laQyeV7WLNBApcR9xC7W"
        }
      ]
    }
    </pre> 
   3) in project directory build the app and run it.
   <pre>
   $ cd boincController
   $ go fmt && golint && go build && ./boincController --config {your config.json}
   </pre>
   
   
   Container available on Dockerhub.com
   <pre>
   $ docker pull pwhite00/boinc_client
   </pre>
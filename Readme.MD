# Dockspawn


Basic Dockerfiles added

# Commands to run VSCode server with python extensions installed

`docker build -t pythonvscode:v0.1 -f PythonDockerfile .`

`docker run -it --init -p 3000:3000 -v "$(pwd):/home/workspace:cached" pythonvscode:v0.1`
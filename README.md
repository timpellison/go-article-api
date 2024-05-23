# Go Service
This repo reflects some of my recent learnings with respect to writing services in Go.  I borrowed deeply from [Mat Ryer's](https://grafana.com/author/mat_ryer/) article "[How I write Go services after 13 years.](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/)".

## Clean Code
One of my colleagues [Kenny Mclive](https://www.linkedin.com/in/kenny-mcclive/) also recently turned me on to
layering techniques to avoid the dreaded circular dependencies with packages.  This concept isn't alien for folks moving 
from .NET or Java and I'm sure we've all fallen into the pit before. These two articles talk a little more about layering approaches so Thank you Kenny!

[Standard Package Layout](https://www.gobeyond.dev/standard-package-layout/)
[Packages as Layers, not Groups](https://www.gobeyond.dev/packages-as-layers/)

## Packages Used
### Gorilla Mux (Server/Routing)
Pretty awesome routing library.  While Mat's article drove me to adopt mux this go round, what I wanted was something that got me closer
to .NET minimal apis and specifically being able to pattern match (e.g., "/api/v1/articles/{id}").  Mux as it turns out readily supports this capability (and much more).

### Viper (Configuration)
This was a choice driven by my experiences with .NET and specifically using static files with environment variable overlays and configuration sections emitted as an object.  

### Gorm (Persistence)

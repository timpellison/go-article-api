[
  {
    "name": "goservice-app",
    "image": "${app_image}",
    "cpu": ${fargate_cpu},
    "memory": ${fargate_memory},
    "networkMode": "awsvpc",
    "environment" : [
        {
            "name" : "DATABASE_CLUSTER",
            "value" : "${database_cluster}"
        },
        {
            "name" : "DATABASE_PASSWORD",
            "value" : "${database_password}"
        },
        {
            "name" : "DATABASE_PORT",
            "value" : "${database_port}"
        },
        {
            "name" : "DATABASE_SSLMODE",
            "value" : "${database_sslmode}"
        }
    ],
    "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "/ecs/cb-app",
          "awslogs-region": "${aws_region}",
          "awslogs-stream-prefix": "ecs"
        }
    },
    "portMappings": [
      {
        "containerPort": ${app_port},
        "hostPort": ${app_port}
      }
    ]
  }
]
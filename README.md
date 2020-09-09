# Repositório GitHub : https://github.com/robsoninocencio/go-hpa.git

# Imagem do app no dockerhub: https://hub.docker.com/repository/docker/robsoni/go-hpa

# serviço do tipo Load Balancer: http://35.202.161.69:8000/

# Para aumentar a carga de acesso ao serviço do tipo LoadBalancer http://35.202.161.69:8000/ siga os passos abaixo:
```
 kubectl exec busybox -- wget -q -O- http://35.202.161.69:8000/
 kubectl exec -it busybox /bin/sh
 while true; do wget -q -O- http://35.202.161.69:8000/; done;
```

# Kubernetes e hpa
## 1) Nessa fase, você deverá implementar exatamente o mesmo algorítimo que foi implementado em PHP (no php-apache-hpa / looping somando a raiz quadrada), porém na linguagem Go Lang. Não esqueça de seguir os passos abaixo:

Nome do deployment, service e nome da imagem deverá se chamar: go-hpa
Desenvolva os testes
Implemente o processo de CI
Faça o push da imagem no Docker Hub (seu-user/go-hpa)
Faça o deploy da imagem no K8S (criando os arquivos de deployment e services)
Cada réplica deverá consumir no mínimo 50m e no máximo 100m.

## 2) Implemente o "hpa" para que esse deployment tenha as seguintes características:

O processo de escala inicia quando a CPU passar de 15%
Quantidade mínima de pods: 1
Quantidade máxima de pods: 6

## 3) Crie um POD e faça requisições através de um looping infinito e verifique se o autoscaler está funcionando corretamente.

## Formato de Entrega:

Crie um repositório no github contendo tanto o projeto em Go (com pelo menos uma PR verificada através do processo de CI), bem como os arquivos yamls necessários para executar todo o processo.

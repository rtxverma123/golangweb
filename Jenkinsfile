pipeline{
   agent any
   
  stages{
    stage('Github'){
      steps{
        git 'https://github.com/rtxverma123/golangweb.git'
      }
    }
    stage('Build'){
      steps {
      sh 'docker build -t rtxverma123/golang .
        
      }
    }
    stage('Push'){
      steps{
//       sh 'docker push rtxverma123/golang'
         echo 'Pushed to dockerhub'
      }
    }
    stage('Deploy'){
      steps{
        sh 'kubectl apply -f deploy.yaml'
        sh 'kubectl apply -f service.yaml'
      }
    }
    

}

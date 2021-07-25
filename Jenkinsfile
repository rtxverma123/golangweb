pipeline{
   agent any
   
  stages{
    stage('Github'){
      steps{
        git ''
      }
    }
    stage('Build'){
      steps {
      sh 'docker build -t rtxverma123/golang .
        
      }
    }
    stage('Push'){
      steps{
      sh 'docker push rtxverma123/golang'
      }
    }
    stage('Deploy'){
      steps{
        sh 'kubectl apply -f deploy.yaml'
        sh 'kubectl apply -f service.yaml'
      }
    }
    

}

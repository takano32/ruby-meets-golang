
task default: 'test'

task 'test' => 'hello.so' do
  sh 'ruby -r./hello -e ""'
end

file 'hello.so' => 'hello.go' do
  sh "GOROOT=#{ENV['PWD']}/go ./go/bin/go build -buildmode=c-shared -o hello.so hello.go"
end

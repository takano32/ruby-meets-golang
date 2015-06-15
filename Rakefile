
task default: 'test'

task 'test' => 'hello.so' do
  sh 'ruby -r./hello -e ""'
end

file 'hello.so' => ['./go/bin/go', 'hello.go'] do
  sh "GOROOT=#{ENV['PWD']}/go ./go/bin/go build -buildmode=c-shared -o hello.so hello.go"
end

file './go/bin/go' do
  #sh 'git clone https://go.googlesource.com/go'
  sh 'cd go/src && ./all.bash'
end

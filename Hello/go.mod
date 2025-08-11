module example.com/Hello

go 1.24.6

replace example.com/Greetings => ../Greetings

require example.com/Greetings v0.0.0-00010101000000-000000000000 // indirect

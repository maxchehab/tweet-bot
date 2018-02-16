parse:
	go run parser.go
train:
	python scripts/preprocess.py --input_txt data/tweet-bot/tweet-bot.txt --output_h5 data/tweet-bot.h5 --output_json data/tweet-bot.json
	th train.lua -input_h5 data/tweet-bot.h5 -input_json data/tweet-bot.json -checkpoint_name output/checkpoint -rnn_size 512 -num_layers 2 -dropout 0.5 -gpu -1

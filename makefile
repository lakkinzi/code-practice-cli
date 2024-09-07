
main := main.go

run:
	go run $(main) -mode=run

create:
	go run $(main) -mode=create

select:
	go run $(main) -mode=select

last_task:
	go run $(main) -mode=last

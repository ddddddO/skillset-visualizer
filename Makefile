run:
	cd app && npm run dev

open:
	cmd.exe /c start http://localhost:8080/#/bar-chart

connpg:
	psql -U postgres -h localhost -d skillsets -p 5432

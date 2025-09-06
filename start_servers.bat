@echo off
echo Starting Student Rooms Backend with WebSocket Support...
echo.

echo Starting GoFr API Server (Port 8000)...
start "API Server" cmd /k "go run cmd/server/main.go"

echo Waiting 3 seconds for API server to start...
timeout /t 3 /nobreak > nul

echo Starting WebSocket Server (Port 8001)...
start "WebSocket Server" cmd /k "go run cmd/websocket/main.go"

echo.
echo Both servers are starting...
echo API Server: http://localhost:8000
echo WebSocket Server: ws://localhost:8001/ws
echo WebSocket Test Client: http://localhost:8000/static/websocket_client.html
echo.
echo Press any key to stop both servers...
pause > nul

echo Stopping servers...
taskkill /f /im go.exe > nul 2>&1
echo Servers stopped.

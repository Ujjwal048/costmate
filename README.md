# CostMate

CostMate is a terminal-based AWS cost monitoring tool that provides a user-friendly interface to track and analyze your AWS service costs. Built with Go and tview, it offers real-time cost data visualization and management capabilities.

## Features

- 📊 Real-time AWS cost monitoring
- 🔄 Multiple AWS profile support
- 💱 Currency conversion (USD/INR)
- 📅 Monthly cost filtering
- 📈 Cost sorting and analysis
- 🖥️ Terminal-based UI with keyboard navigation

## Prerequisites

- Go 1.22 or higher
- AWS CLI configured with profiles
- AWS credentials with Cost Explorer permissions

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/costmate.git
cd costmate
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o costmate cmd/app/main.go
```

## Usage

Run the application:
```bash
./costmate
```

### Keyboard Controls

- `p` - Switch AWS Profile
- `c` - Toggle between USD and INR currencies
- `s` - Sort services by cost
- `m` - Filter costs by month
- `↑/↓` - Navigate through services
- `Esc` - Close modals/return to main view

## Configuration

### AWS Profiles
CostMate uses your AWS CLI profiles. Configure them in `~/.aws/credentials`:
```ini
[default]
aws_access_key_id = YOUR_ACCESS_KEY
aws_secret_access_key = YOUR_SECRET_KEY

[profile2]
aws_access_key_id = YOUR_ACCESS_KEY
aws_secret_access_key = YOUR_SECRET_KEY
```

### Currency Conversion
The application automatically fetches current USD to INR conversion rates from the Frankfurter API.

## Project Structure

```
costmate/
├── cmd/
│   └── app/
│       └── main.go           # Application entry point
├── internal/
│   ├── aws/                 # AWS SDK integration
│   ├── bootstrap/           # Application initialization
│   ├── constants/           # Application constants
│   ├── logger/             # Logging functionality
│   ├── modals/             # UI modal components
│   ├── ui/                 # UI components
│   └── utils/              # Utility functions
└── go.mod                  # Go module file
```

## Development

### Adding New Features
1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

### Running Tests
```bash
go test ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [tview](https://github.com/rivo/tview) for the terminal UI framework
- [AWS SDK for Go](https://github.com/aws/aws-sdk-go-v2) for AWS integration
- [Frankfurter API](https://www.frankfurter.app/) for currency conversion rates 
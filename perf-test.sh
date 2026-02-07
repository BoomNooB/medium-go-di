#!/bin/bash

# Performance comparison script for DI changes
# Usage: ./perf-test.sh [before|after]

set -e

if [ "$1" != "before" ] && [ "$1" != "after" ]; then
    echo "Usage: ./perf-test.sh [before|after]"
    echo ""
    echo "Examples:"
    echo "  ./perf-test.sh before   # Run baseline test before DI changes"
    echo "  ./perf-test.sh after    # Run test after DI changes"
    exit 1
fi

STAGE="$1"
TIMESTAMP=$(date +%Y%m%d-%H%M%S)
OUTPUT_DIR="k6/results"
JSON_FILE="${OUTPUT_DIR}/results-${STAGE}-di-${TIMESTAMP}.json"
TXT_FILE="${OUTPUT_DIR}/results-${STAGE}-di-${TIMESTAMP}.txt"

# Create results directory
mkdir -p "$OUTPUT_DIR"

echo "================================================"
echo "Performance Test: ${STAGE} DI changes"
echo "================================================"
echo ""

# Stop any running containers
echo "📦 Stopping any running containers..."
docker-compose down 2>/dev/null || true

# Build and start the application
echo "🏗️  Building and starting application in Docker..."
echo "   (Limited to 64MB RAM, 0.25 CPU)"
docker-compose up --build -d

# Wait for the application to be ready
echo "⏳ Waiting for application to be ready..."
sleep 3

# Check if the application is responding
if ! curl -s http://localhost:1323/api/v1/favorite > /dev/null 2>&1; then
    echo "⚠️  Warning: Application might not be ready yet. Waiting a bit more..."
    sleep 3
fi

echo "✅ Application is ready!"
echo ""

# Show resource limits
echo "📊 Container resource limits:"
docker stats --no-stream --format "table {{.Container}}\t{{.CPUPerc}}\t{{.MemUsage}}\t{{.MemPerc}}" $(docker-compose ps -q)
echo ""

# Run the load test
echo "🚀 Starting k6 load test..."
echo "   Results will be saved to:"
echo "   - ${JSON_FILE}"
echo "   - ${TXT_FILE}"
echo ""

cd k6
k6 run loadtest.js --out json="../${JSON_FILE}" | tee "../${TXT_FILE}"
cd ..

echo ""
echo "================================================"
echo "✅ Test completed!"
echo "================================================"
echo ""
echo "📈 Summary saved to: ${TXT_FILE}"
echo ""

# Extract key metrics from the text file
echo "Key Metrics:"
echo "------------"
grep -E "(http_req_duration|http_req_failed|http_reqs[^_]|checks)" "${TXT_FILE}" | sed 's/^/  /'
echo ""

# Show container resource usage during test
echo "📊 Final resource usage:"
docker stats --no-stream --format "table {{.Container}}\t{{.CPUPerc}}\t{{.MemUsage}}\t{{.MemPerc}}" $(docker-compose ps -q)
echo ""

# Keep container running for manual inspection
echo "🔧 Container is still running for inspection."
echo "   To view logs: docker-compose logs -f"
echo "   To stop: docker-compose down"
echo ""

if [ "$STAGE" == "before" ]; then
    echo "💡 Next steps:"
    echo "   1. Make your DI changes"
    echo "   2. Run: ./perf-test.sh after"
    echo "   3. Compare the results files"
fi

if [ "$STAGE" == "after" ]; then
    echo "💡 Compare results:"
    echo "   Check the results files in ${OUTPUT_DIR}/"
    echo "   Look for differences in:"
    echo "   - http_req_duration (p95, p99)"
    echo "   - http_reqs (requests per second)"
    echo "   - http_req_failed (error rate)"
fi

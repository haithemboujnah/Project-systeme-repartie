import grpc
from concurrent import futures
import time
import factorization_pb2
import factorization_pb2_grpc

def prime_factors(n):
    factors = []
    divisor = 2
    while divisor <= n:
        if n % divisor == 0:
            factors.append(divisor)
            n //= divisor
        else:
            divisor += 1
    return factors

class PrimeFactorizationService(factorization_pb2_grpc.PrimeFactorizationServiceServicer):
    def GetPrimeFactors(self, request, context):
        factors = prime_factors(request.number)
        return factorization_pb2.PrimeFactorsResponse(factors=factors)
    
    def GetPrimeFactorsStream(self, request_iterator, context):
        for request in request_iterator:
            factors = prime_factors(request.number)
            yield factorization_pb2.PrimeFactorsResponse(factors=factors)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    factorization_pb2_grpc.add_PrimeFactorizationServiceServicer_to_server(PrimeFactorizationService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("Server started, listening on port 50051")
    try:
        while True:
            time.sleep(86400)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    serve()

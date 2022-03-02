import time

'''
These tests are meant to verify the fa_client actually caches the results of each query. 
The 2 seconds sleep between two calls to the same method should be enough to guarantee
that if those were two differen calls to the FA endpoint the sampled metrics are different.
'''

def test_arrays(fa_client):
    arrays1 = fa_client.arrays()
    time.sleep(2)
    arrays2 = fa_client.arrays()
    print(arrays1)
    assert arrays1 == arrays2

# you are required to install -  mmh3, bitarray 

import math
import mmh3
from bitarray import bitarray


class BloomFilter(object):

	def __init__(self, items_count, fp_prob):

		# items_count -> number of iterms to be stored -> int
		# fb_prob -> false positive probability -> float

		self.fp_prob = fp_prob

		#  first calculte the size of bit array  and initialise the bit array with 0 value
		#  second, find the number of hash functions
		 
		# setting the size of bit array to use 
		self.size = self.get_size(items_count, fp_prob) 

		# s the value of bitarray to zero
		self.bit_array = bitarray(self.size)

		# initializing all bits as 0
		self.bit_array.setall(0)

		# getting the number of hash functions to use
		self.hash_count = self.get_hash_count(self.size, items_count)



#  function to add an item in the bit array
	def add(self,item):

		digests = []

		for i in range(self.hash_count):
			
			
			# i here serves as the seed
			digest = mmh3.hash(item,i) % self.size
			digests.append(digest)

			self.bit_array[digest] = True

#  function to check if an element exist in the array or not 
	def check(self,item):

		for i in range( self.hash_count):
			digest = mmh3.hash(item,i) % self.size 

			if self.bit_array[digest] == False :
				# basic logic -> if any bit is false, then it is not present in th filter... else there is probablility it is present

				return False
			
		return True

# 	defining class methods 

	@classmethod
	def get_size(self,n,p):

		m = -(n * math.log(p))/(math.log(2)**2)
		return int(m)

	@classmethod
	def get_hash_count(self, m, n):
		k = (m/n) * math.log(2)
		return int(k)

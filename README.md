# Wirecutter

This is a choas engineering project whose purpose is to degrade kubernetes 
services without impating pod execution (no kill/eviction of pods) by woking
only on the services selectors and related labels in the pods.

To achieve this goal Wirecutter must keep the original state of the altered 
resource and roll back it after a certain amount of time.

# Maintainers
Gianni Salinetti <gsalinet@redhat.com>


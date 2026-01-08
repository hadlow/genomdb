# GenomDB

GenomDB is a distributed storage system for genomic data. GenomDB can store both SAM and BAM files. Files are sharded and compressed.

## How it works

Each read of the SAM file will be split across multiple nodes in the network.


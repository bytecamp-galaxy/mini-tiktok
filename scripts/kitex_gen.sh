#!/usr/bin/env bash
kitex --module github.com/bytecamp-galaxy/mini-tiktok ../idl/rpc/comment_service.thrift
kitex --module github.com/bytecamp-galaxy/mini-tiktok ../idl/rpc/favorite_service.thrift
kitex --module github.com/bytecamp-galaxy/mini-tiktok ../idl/rpc/feed_service.thrift
kitex --module github.com/bytecamp-galaxy/mini-tiktok ../idl/rpc/publish_service.thrift
kitex --module github.com/bytecamp-galaxy/mini-tiktok ../idl/rpc/user_service.thrift

#rm -rf ../kitex_gen/feed
#mv kitex_gen/feed ../kitex_gen

#rm -rf ../kitex_gen/publish
#mv kitex_gen/publish ../kitex_gen
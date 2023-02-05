#!/usr/bin/env bash
# set working directory to project root
kitex --module github.com/bytecamp-galaxy/mini-tiktok idl/rpc/comment_service.thrift
kitex --module github.com/bytecamp-galaxy/mini-tiktok idl/rpc/favorite_service.thrift
kitex --module github.com/bytecamp-galaxy/mini-tiktok idl/rpc/feed_service.thrift
kitex --module github.com/bytecamp-galaxy/mini-tiktok idl/rpc/publish_service.thrift
kitex --module github.com/bytecamp-galaxy/mini-tiktok idl/rpc/user_service.thrift
kitex --module github.com/bytecamp-galaxy/mini-tiktok idl/rpc/relation_service.thrift
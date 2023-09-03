// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by cmd/codegen. DO NOT EDIT.
//
// Thumbnail Option Group

package ytdlp

// Write thumbnail image to disk
//
//   - See also [UnsetWriteThumbnail], for unsetting the flag.
//   - WriteThumbnail maps to cli flags: --write-thumbnail.
func (c *Command) WriteThumbnail(value string) *Command {
	c.addFlag(&Flag{
		ID:   "writethumbnail",
		Flag: "--write-thumbnail",
		Args: []string{value},
	})
	return c
}

// UnsetWriteThumbnail unsets any flags that were previously set by
// [WriteThumbnail].
func (c *Command) UnsetWriteThumbnail() *Command {
	c.removeFlagByID("writethumbnail")
	return c
}

// Do not write thumbnail image to disk (default)
//
//   - See also [UnsetNoWriteThumbnail], for unsetting the flag.
//   - NoWriteThumbnail maps to cli flags: --no-write-thumbnail.
func (c *Command) NoWriteThumbnail() *Command {
	c.addFlag(&Flag{
		ID:   "writethumbnail",
		Flag: "--no-write-thumbnail",
		Args: nil,
	})
	return c
}

// UnsetNoWriteThumbnail unsets any flags that were previously set by
// [NoWriteThumbnail].
func (c *Command) UnsetNoWriteThumbnail() *Command {
	c.removeFlagByID("writethumbnail")
	return c
}

// List available thumbnails of each video. Simulate unless --no-simulate is used
//
//   - See also [UnsetListThumbnails], for unsetting the flag.
//   - ListThumbnails maps to cli flags: --list-thumbnails.
func (c *Command) ListThumbnails() *Command {
	c.addFlag(&Flag{
		ID:   "list_thumbnails",
		Flag: "--list-thumbnails",
		Args: nil,
	})
	return c
}

// UnsetListThumbnails unsets any flags that were previously set by
// [ListThumbnails].
func (c *Command) UnsetListThumbnails() *Command {
	c.removeFlagByID("list_thumbnails")
	return c
}

// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by cmd/codegen. DO NOT EDIT.
//
// Download Option Group

package ytdlp

import (
	"strconv"
)

// Number of fragments of a dash/hlsnative video that should be downloaded
// concurrently (default is 1)
//
//   - See also [UnsetConcurrentFragments], for unsetting the flag.
//   - ConcurrentFragments maps to cli flags: -N/--concurrent-fragments=N.
func (c *Command) ConcurrentFragments(n int) *Command {
	c.addFlag(&Flag{
		ID:   "concurrent_fragment_downloads",
		Flag: "--concurrent-fragments",
		Args: []string{
			strconv.Itoa(n),
		},
	})
	return c
}

// UnsetConcurrentFragments unsets any flags that were previously set by
// [ConcurrentFragments].
func (c *Command) UnsetConcurrentFragments() *Command {
	c.removeFlagByID("concurrent_fragment_downloads")
	return c
}

// Maximum download rate in bytes per second, e.g. 50K or 4.2M
//
//   - See also [UnsetLimitRate], for unsetting the flag.
//   - LimitRate maps to cli flags: -r/--limit-rate/--rate-limit=RATE.
func (c *Command) LimitRate(rate string) *Command {
	c.addFlag(&Flag{
		ID:   "ratelimit",
		Flag: "--limit-rate",
		Args: []string{rate},
	})
	return c
}

// UnsetLimitRate unsets any flags that were previously set by
// [LimitRate].
func (c *Command) UnsetLimitRate() *Command {
	c.removeFlagByID("ratelimit")
	return c
}

// Minimum download rate in bytes per second below which throttling is assumed and
// the video data is re-extracted, e.g. 100K
//
//   - See also [UnsetThrottledRate], for unsetting the flag.
//   - ThrottledRate maps to cli flags: --throttled-rate=RATE.
func (c *Command) ThrottledRate(rate string) *Command {
	c.addFlag(&Flag{
		ID:   "throttledratelimit",
		Flag: "--throttled-rate",
		Args: []string{rate},
	})
	return c
}

// UnsetThrottledRate unsets any flags that were previously set by
// [ThrottledRate].
func (c *Command) UnsetThrottledRate() *Command {
	c.removeFlagByID("throttledratelimit")
	return c
}

// Number of retries (default is 10), or "infinite"
//
//   - See also [UnsetRetries], for unsetting the flag.
//   - Retries maps to cli flags: -R/--retries=RETRIES.
func (c *Command) Retries(retries string) *Command {
	c.addFlag(&Flag{
		ID:   "retries",
		Flag: "--retries",
		Args: []string{retries},
	})
	return c
}

// UnsetRetries unsets any flags that were previously set by
// [Retries].
func (c *Command) UnsetRetries() *Command {
	c.removeFlagByID("retries")
	return c
}

// Number of times to retry on file access error (default is 3), or "infinite"
//
//   - See also [UnsetFileAccessRetries], for unsetting the flag.
//   - FileAccessRetries maps to cli flags: --file-access-retries=RETRIES.
func (c *Command) FileAccessRetries(retries string) *Command {
	c.addFlag(&Flag{
		ID:   "file_access_retries",
		Flag: "--file-access-retries",
		Args: []string{retries},
	})
	return c
}

// UnsetFileAccessRetries unsets any flags that were previously set by
// [FileAccessRetries].
func (c *Command) UnsetFileAccessRetries() *Command {
	c.removeFlagByID("file_access_retries")
	return c
}

// Number of retries for a fragment (default is 10), or "infinite" (DASH, hlsnative
// and ISM)
//
//   - See also [UnsetFragmentRetries], for unsetting the flag.
//   - FragmentRetries maps to cli flags: --fragment-retries=RETRIES.
func (c *Command) FragmentRetries(retries string) *Command {
	c.addFlag(&Flag{
		ID:   "fragment_retries",
		Flag: "--fragment-retries",
		Args: []string{retries},
	})
	return c
}

// UnsetFragmentRetries unsets any flags that were previously set by
// [FragmentRetries].
func (c *Command) UnsetFragmentRetries() *Command {
	c.removeFlagByID("fragment_retries")
	return c
}

// Time to sleep between retries in seconds (optionally) prefixed by the type of
// retry (http (default), fragment, file_access, extractor) to apply the sleep to.
// EXPR can be a number, linear=START[:END[:STEP=1]] or exp=START[:END[:BASE=2]].
// This option can be used multiple times to set the sleep for the different retry
// types, e.g. --retry-sleep linear=1::2 --retry-sleep fragment:exp=1:20
//
//   - See also [UnsetRetrySleep], for unsetting the flag.
//   - RetrySleep maps to cli flags: --retry-sleep=[TYPE:]EXPR.
func (c *Command) RetrySleep(expr string) *Command {
	c.addFlag(&Flag{
		ID:   "retry_sleep",
		Flag: "--retry-sleep",
		Args: []string{expr},
	})
	return c
}

// UnsetRetrySleep unsets any flags that were previously set by
// [RetrySleep].
func (c *Command) UnsetRetrySleep() *Command {
	c.removeFlagByID("retry_sleep")
	return c
}

// Skip unavailable fragments for DASH, hlsnative and ISM downloads (default)
// (Alias: --no-abort-on-unavailable-fragments)
//
//   - See also [UnsetSkipUnavailableFragments], for unsetting the flag.
//   - SkipUnavailableFragments maps to cli flags: --skip-unavailable-fragments/--no-abort-on-unavailable-fragments.
func (c *Command) SkipUnavailableFragments() *Command {
	c.addFlag(&Flag{
		ID:   "skip_unavailable_fragments",
		Flag: "--skip-unavailable-fragments",
		Args: nil,
	})
	return c
}

// UnsetSkipUnavailableFragments unsets any flags that were previously set by
// [SkipUnavailableFragments].
func (c *Command) UnsetSkipUnavailableFragments() *Command {
	c.removeFlagByID("skip_unavailable_fragments")
	return c
}

// Abort download if a fragment is unavailable (Alias:
// --no-skip-unavailable-fragments)
//
//   - See also [UnsetAbortOnUnavailableFragments], for unsetting the flag.
//   - AbortOnUnavailableFragments maps to cli flags: --abort-on-unavailable-fragments/--no-skip-unavailable-fragments.
func (c *Command) AbortOnUnavailableFragments() *Command {
	c.addFlag(&Flag{
		ID:   "skip_unavailable_fragments",
		Flag: "--abort-on-unavailable-fragments",
		Args: nil,
	})
	return c
}

// UnsetAbortOnUnavailableFragments unsets any flags that were previously set by
// [AbortOnUnavailableFragments].
func (c *Command) UnsetAbortOnUnavailableFragments() *Command {
	c.removeFlagByID("skip_unavailable_fragments")
	return c
}

// Keep downloaded fragments on disk after downloading is finished
//
//   - See also [UnsetKeepFragments], for unsetting the flag.
//   - KeepFragments maps to cli flags: --keep-fragments.
func (c *Command) KeepFragments() *Command {
	c.addFlag(&Flag{
		ID:   "keep_fragments",
		Flag: "--keep-fragments",
		Args: nil,
	})
	return c
}

// UnsetKeepFragments unsets any flags that were previously set by
// [KeepFragments].
func (c *Command) UnsetKeepFragments() *Command {
	c.removeFlagByID("keep_fragments")
	return c
}

// Delete downloaded fragments after downloading is finished (default)
//
//   - See also [UnsetNoKeepFragments], for unsetting the flag.
//   - NoKeepFragments maps to cli flags: --no-keep-fragments.
func (c *Command) NoKeepFragments() *Command {
	c.addFlag(&Flag{
		ID:   "keep_fragments",
		Flag: "--no-keep-fragments",
		Args: nil,
	})
	return c
}

// UnsetNoKeepFragments unsets any flags that were previously set by
// [NoKeepFragments].
func (c *Command) UnsetNoKeepFragments() *Command {
	c.removeFlagByID("keep_fragments")
	return c
}

// Size of download buffer, e.g. 1024 or 16K (default is 1024)
//
//   - See also [UnsetBufferSize], for unsetting the flag.
//   - BufferSize maps to cli flags: --buffer-size=SIZE.
func (c *Command) BufferSize(size string) *Command {
	c.addFlag(&Flag{
		ID:   "buffersize",
		Flag: "--buffer-size",
		Args: []string{size},
	})
	return c
}

// UnsetBufferSize unsets any flags that were previously set by
// [BufferSize].
func (c *Command) UnsetBufferSize() *Command {
	c.removeFlagByID("buffersize")
	return c
}

// The buffer size is automatically resized from an initial value of --buffer-size
// (default)
//
//   - See also [UnsetResizeBuffer], for unsetting the flag.
//   - ResizeBuffer maps to cli flags: --resize-buffer.
func (c *Command) ResizeBuffer() *Command {
	c.addFlag(&Flag{
		ID:   "noresizebuffer",
		Flag: "--resize-buffer",
		Args: nil,
	})
	return c
}

// UnsetResizeBuffer unsets any flags that were previously set by
// [ResizeBuffer].
func (c *Command) UnsetResizeBuffer() *Command {
	c.removeFlagByID("noresizebuffer")
	return c
}

// Do not automatically adjust the buffer size
//
//   - See also [UnsetNoResizeBuffer], for unsetting the flag.
//   - NoResizeBuffer maps to cli flags: --no-resize-buffer.
func (c *Command) NoResizeBuffer() *Command {
	c.addFlag(&Flag{
		ID:   "noresizebuffer",
		Flag: "--no-resize-buffer",
		Args: nil,
	})
	return c
}

// UnsetNoResizeBuffer unsets any flags that were previously set by
// [NoResizeBuffer].
func (c *Command) UnsetNoResizeBuffer() *Command {
	c.removeFlagByID("noresizebuffer")
	return c
}

// Size of a chunk for chunk-based HTTP downloading, e.g. 10485760 or 10M (default
// is disabled). May be useful for bypassing bandwidth throttling imposed by a
// webserver (experimental)
//
//   - See also [UnsetHttpChunkSize], for unsetting the flag.
//   - HttpChunkSize maps to cli flags: --http-chunk-size=SIZE.
func (c *Command) HttpChunkSize(size string) *Command {
	c.addFlag(&Flag{
		ID:   "http_chunk_size",
		Flag: "--http-chunk-size",
		Args: []string{size},
	})
	return c
}

// UnsetHttpChunkSize unsets any flags that were previously set by
// [HttpChunkSize].
func (c *Command) UnsetHttpChunkSize() *Command {
	c.removeFlagByID("http_chunk_size")
	return c
}

// Test sets the "test" flag (no description specified).
//
//   - See also [UnsetTest], for unsetting the flag.
//   - Test maps to cli flags: --test (hidden).
func (c *Command) Test() *Command {
	c.addFlag(&Flag{
		ID:   "test",
		Flag: "--test",
		Args: nil,
	})
	return c
}

// UnsetTest unsets any flags that were previously set by
// [Test].
func (c *Command) UnsetTest() *Command {
	c.removeFlagByID("test")
	return c
}

// PlaylistReverse sets the "playlist-reverse" flag (no description specified).
//
//   - See also [UnsetPlaylistReverse], for unsetting the flag.
//   - PlaylistReverse maps to cli flags: --playlist-reverse (hidden).
func (c *Command) PlaylistReverse() *Command {
	c.addFlag(&Flag{
		ID:   "playlist_reverse",
		Flag: "--playlist-reverse",
		Args: nil,
	})
	return c
}

// UnsetPlaylistReverse unsets any flags that were previously set by
// [PlaylistReverse].
func (c *Command) UnsetPlaylistReverse() *Command {
	c.removeFlagByID("playlist_reverse")
	return c
}

// NoPlaylistReverse sets the "no-playlist-reverse" flag (no description specified).
//
//   - See also [UnsetNoPlaylistReverse], for unsetting the flag.
//   - NoPlaylistReverse maps to cli flags: --no-playlist-reverse (hidden).
func (c *Command) NoPlaylistReverse() *Command {
	c.addFlag(&Flag{
		ID:   "playlist_reverse",
		Flag: "--no-playlist-reverse",
		Args: nil,
	})
	return c
}

// UnsetNoPlaylistReverse unsets any flags that were previously set by
// [NoPlaylistReverse].
func (c *Command) UnsetNoPlaylistReverse() *Command {
	c.removeFlagByID("playlist_reverse")
	return c
}

// Download playlist videos in random order
//
//   - See also [UnsetPlaylistRandom], for unsetting the flag.
//   - PlaylistRandom maps to cli flags: --playlist-random.
func (c *Command) PlaylistRandom() *Command {
	c.addFlag(&Flag{
		ID:   "playlist_random",
		Flag: "--playlist-random",
		Args: nil,
	})
	return c
}

// UnsetPlaylistRandom unsets any flags that were previously set by
// [PlaylistRandom].
func (c *Command) UnsetPlaylistRandom() *Command {
	c.removeFlagByID("playlist_random")
	return c
}

// Process entries in the playlist as they are received. This disables n_entries,
// --playlist-random and --playlist-reverse
//
//   - See also [UnsetLazyPlaylist], for unsetting the flag.
//   - LazyPlaylist maps to cli flags: --lazy-playlist.
func (c *Command) LazyPlaylist() *Command {
	c.addFlag(&Flag{
		ID:   "lazy_playlist",
		Flag: "--lazy-playlist",
		Args: nil,
	})
	return c
}

// UnsetLazyPlaylist unsets any flags that were previously set by
// [LazyPlaylist].
func (c *Command) UnsetLazyPlaylist() *Command {
	c.removeFlagByID("lazy_playlist")
	return c
}

// Process videos in the playlist only after the entire playlist is parsed
// (default)
//
//   - See also [UnsetNoLazyPlaylist], for unsetting the flag.
//   - NoLazyPlaylist maps to cli flags: --no-lazy-playlist.
func (c *Command) NoLazyPlaylist() *Command {
	c.addFlag(&Flag{
		ID:   "lazy_playlist",
		Flag: "--no-lazy-playlist",
		Args: nil,
	})
	return c
}

// UnsetNoLazyPlaylist unsets any flags that were previously set by
// [NoLazyPlaylist].
func (c *Command) UnsetNoLazyPlaylist() *Command {
	c.removeFlagByID("lazy_playlist")
	return c
}

// Set file xattribute ytdl.filesize with expected file size
//
//   - See also [UnsetXattrSetFilesize], for unsetting the flag.
//   - XattrSetFilesize maps to cli flags: --xattr-set-filesize.
func (c *Command) XattrSetFilesize() *Command {
	c.addFlag(&Flag{
		ID:   "xattr_set_filesize",
		Flag: "--xattr-set-filesize",
		Args: nil,
	})
	return c
}

// UnsetXattrSetFilesize unsets any flags that were previously set by
// [XattrSetFilesize].
func (c *Command) UnsetXattrSetFilesize() *Command {
	c.removeFlagByID("xattr_set_filesize")
	return c
}

// HlsPreferNative sets the "hls-prefer-native" flag (no description specified).
//
//   - See also [UnsetHlsPreferNative], for unsetting the flag.
//   - HlsPreferNative maps to cli flags: --hls-prefer-native (hidden).
func (c *Command) HlsPreferNative() *Command {
	c.addFlag(&Flag{
		ID:   "hls_prefer_native",
		Flag: "--hls-prefer-native",
		Args: nil,
	})
	return c
}

// UnsetHlsPreferNative unsets any flags that were previously set by
// [HlsPreferNative].
func (c *Command) UnsetHlsPreferNative() *Command {
	c.removeFlagByID("hls_prefer_native")
	return c
}

// HlsPreferFfmpeg sets the "hls-prefer-ffmpeg" flag (no description specified).
//
//   - See also [UnsetHlsPreferFfmpeg], for unsetting the flag.
//   - HlsPreferFfmpeg maps to cli flags: --hls-prefer-ffmpeg (hidden).
func (c *Command) HlsPreferFfmpeg() *Command {
	c.addFlag(&Flag{
		ID:   "hls_prefer_native",
		Flag: "--hls-prefer-ffmpeg",
		Args: nil,
	})
	return c
}

// UnsetHlsPreferFfmpeg unsets any flags that were previously set by
// [HlsPreferFfmpeg].
func (c *Command) UnsetHlsPreferFfmpeg() *Command {
	c.removeFlagByID("hls_prefer_native")
	return c
}

// Use the mpegts container for HLS videos; allowing some players to play the video
// while downloading, and reducing the chance of file corruption if download is
// interrupted. This is enabled by default for live streams
//
//   - See also [UnsetHlsUseMpegts], for unsetting the flag.
//   - HlsUseMpegts maps to cli flags: --hls-use-mpegts.
func (c *Command) HlsUseMpegts() *Command {
	c.addFlag(&Flag{
		ID:   "hls_use_mpegts",
		Flag: "--hls-use-mpegts",
		Args: nil,
	})
	return c
}

// UnsetHlsUseMpegts unsets any flags that were previously set by
// [HlsUseMpegts].
func (c *Command) UnsetHlsUseMpegts() *Command {
	c.removeFlagByID("hls_use_mpegts")
	return c
}

// Do not use the mpegts container for HLS videos. This is default when not
// downloading live streams
//
//   - See also [UnsetNoHlsUseMpegts], for unsetting the flag.
//   - NoHlsUseMpegts maps to cli flags: --no-hls-use-mpegts.
func (c *Command) NoHlsUseMpegts() *Command {
	c.addFlag(&Flag{
		ID:   "hls_use_mpegts",
		Flag: "--no-hls-use-mpegts",
		Args: nil,
	})
	return c
}

// UnsetNoHlsUseMpegts unsets any flags that were previously set by
// [NoHlsUseMpegts].
func (c *Command) UnsetNoHlsUseMpegts() *Command {
	c.removeFlagByID("hls_use_mpegts")
	return c
}

// Download only chapters that match the regular expression. A "*" prefix denotes
// time-range instead of chapter. Negative timestamps are calculated from the end.
// "*from-url" can be used to download between the "start_time" and "end_time"
// extracted from the URL. Needs ffmpeg. This option can be used multiple times to
// download multiple sections, e.g. --download-sections "*10:15-inf"
// --download-sections "intro"
//
//   - See also [UnsetDownloadSections], for unsetting the flag.
//   - DownloadSections maps to cli flags: --download-sections=REGEX.
func (c *Command) DownloadSections(regex string) *Command {
	c.addFlag(&Flag{
		ID:   "download_ranges",
		Flag: "--download-sections",
		Args: []string{regex},
	})
	return c
}

// UnsetDownloadSections unsets any flags that were previously set by
// [DownloadSections].
func (c *Command) UnsetDownloadSections() *Command {
	c.removeFlagByID("download_ranges")
	return c
}

// Name or path of the external downloader to use (optionally) prefixed by the
// protocols (http, ftp, m3u8, dash, rstp, rtmp, mms) to use it for. Currently
// supports native, aria2c, avconv, axel, curl, ffmpeg, httpie, wget. You can use
// this option multiple times to set different downloaders for different protocols.
// E.g. --downloader aria2c --downloader "dash,m3u8:native" will use aria2c for
// http/ftp downloads, and the native downloader for dash/m3u8 downloads (Alias:
// --external-downloader)
//
//   - See also [UnsetDownloader], for unsetting the flag.
//   - Downloader maps to cli flags: --downloader/--external-downloader=[PROTO:]NAME.
func (c *Command) Downloader(name string) *Command {
	c.addFlag(&Flag{
		ID:   "external_downloader",
		Flag: "--downloader",
		Args: []string{name},
	})
	return c
}

// UnsetDownloader unsets any flags that were previously set by
// [Downloader].
func (c *Command) UnsetDownloader() *Command {
	c.removeFlagByID("external_downloader")
	return c
}

// Give these arguments to the external downloader. Specify the downloader name and
// the arguments separated by a colon ":". For ffmpeg, arguments can be passed to
// different positions using the same syntax as --postprocessor-args. You can use
// this option multiple times to give different arguments to different downloaders
// (Alias: --external-downloader-args)
//
//   - See also [UnsetDownloaderArgs], for unsetting the flag.
//   - DownloaderArgs maps to cli flags: --downloader-args/--external-downloader-args=NAME:ARGS.
func (c *Command) DownloaderArgs(nameargs string) *Command {
	c.addFlag(&Flag{
		ID:   "external_downloader_args",
		Flag: "--downloader-args",
		Args: []string{nameargs},
	})
	return c
}

// UnsetDownloaderArgs unsets any flags that were previously set by
// [DownloaderArgs].
func (c *Command) UnsetDownloaderArgs() *Command {
	c.removeFlagByID("external_downloader_args")
	return c
}

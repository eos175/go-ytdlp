// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by cmd/codegen. DO NOT EDIT.
//
// Extractor Option Group

package ytdlp

// Number of retries for known extractor errors (default is 3), or "infinite"
//
//   - See also [UnsetExtractorRetries], for unsetting the flag.
//   - ExtractorRetries maps to cli flags: --extractor-retries=RETRIES.
func (c *Command) ExtractorRetries(retries string) *Command {
	c.addFlag(&Flag{
		ID:   "extractor_retries",
		Flag: "--extractor-retries",
		Args: []string{retries},
	})
	return c
}

// UnsetExtractorRetries unsets any flags that were previously set by
// [ExtractorRetries].
func (c *Command) UnsetExtractorRetries() *Command {
	c.removeFlagByID("extractor_retries")
	return c
}

// Process dynamic DASH manifests (default) (Alias: --no-ignore-dynamic-mpd)
//
//   - See also [UnsetAllowDynamicMpd], for unsetting the flag.
//   - AllowDynamicMpd maps to cli flags: --allow-dynamic-mpd/--no-ignore-dynamic-mpd.
func (c *Command) AllowDynamicMpd() *Command {
	c.addFlag(&Flag{
		ID:   "dynamic_mpd",
		Flag: "--allow-dynamic-mpd",
		Args: nil,
	})
	return c
}

// UnsetAllowDynamicMpd unsets any flags that were previously set by
// [AllowDynamicMpd].
func (c *Command) UnsetAllowDynamicMpd() *Command {
	c.removeFlagByID("dynamic_mpd")
	return c
}

// Do not process dynamic DASH manifests (Alias: --no-allow-dynamic-mpd)
//
//   - See also [UnsetIgnoreDynamicMpd], for unsetting the flag.
//   - IgnoreDynamicMpd maps to cli flags: --ignore-dynamic-mpd/--no-allow-dynamic-mpd.
func (c *Command) IgnoreDynamicMpd() *Command {
	c.addFlag(&Flag{
		ID:   "dynamic_mpd",
		Flag: "--ignore-dynamic-mpd",
		Args: nil,
	})
	return c
}

// UnsetIgnoreDynamicMpd unsets any flags that were previously set by
// [IgnoreDynamicMpd].
func (c *Command) UnsetIgnoreDynamicMpd() *Command {
	c.removeFlagByID("dynamic_mpd")
	return c
}

// Split HLS playlists to different formats at discontinuities such as ad breaks
//
//   - See also [UnsetHlsSplitDiscontinuity], for unsetting the flag.
//   - HlsSplitDiscontinuity maps to cli flags: --hls-split-discontinuity.
func (c *Command) HlsSplitDiscontinuity() *Command {
	c.addFlag(&Flag{
		ID:   "hls_split_discontinuity",
		Flag: "--hls-split-discontinuity",
		Args: nil,
	})
	return c
}

// UnsetHlsSplitDiscontinuity unsets any flags that were previously set by
// [HlsSplitDiscontinuity].
func (c *Command) UnsetHlsSplitDiscontinuity() *Command {
	c.removeFlagByID("hls_split_discontinuity")
	return c
}

// Do not split HLS playlists to different formats at discontinuities such as ad
// breaks (default)
//
//   - See also [UnsetNoHlsSplitDiscontinuity], for unsetting the flag.
//   - NoHlsSplitDiscontinuity maps to cli flags: --no-hls-split-discontinuity.
func (c *Command) NoHlsSplitDiscontinuity() *Command {
	c.addFlag(&Flag{
		ID:   "hls_split_discontinuity",
		Flag: "--no-hls-split-discontinuity",
		Args: nil,
	})
	return c
}

// UnsetNoHlsSplitDiscontinuity unsets any flags that were previously set by
// [NoHlsSplitDiscontinuity].
func (c *Command) UnsetNoHlsSplitDiscontinuity() *Command {
	c.removeFlagByID("hls_split_discontinuity")
	return c
}

// Pass ARGS arguments to the IE_KEY extractor. See "EXTRACTOR ARGUMENTS" for
// details. You can use this option multiple times to give arguments for different
// extractors
//
//   - See also [UnsetExtractorArgs], for unsetting the flag.
//   - ExtractorArgs maps to cli flags: --extractor-args=IE_KEY:ARGS.
func (c *Command) ExtractorArgs(ieKeyargs string) *Command {
	c.addFlag(&Flag{
		ID:   "extractor_args",
		Flag: "--extractor-args",
		Args: []string{ieKeyargs},
	})
	return c
}

// UnsetExtractorArgs unsets any flags that were previously set by
// [ExtractorArgs].
func (c *Command) UnsetExtractorArgs() *Command {
	c.removeFlagByID("extractor_args")
	return c
}

// YoutubeIncludeDashManifest sets the "youtube-include-dash-manifest" flag (no description specified).
//
//   - See also [UnsetYoutubeIncludeDashManifest], for unsetting the flag.
//   - YoutubeIncludeDashManifest maps to cli flags: --youtube-include-dash-manifest/--no-youtube-skip-dash-manifest (hidden).
func (c *Command) YoutubeIncludeDashManifest() *Command {
	c.addFlag(&Flag{
		ID:   "youtube_include_dash_manifest",
		Flag: "--youtube-include-dash-manifest",
		Args: nil,
	})
	return c
}

// UnsetYoutubeIncludeDashManifest unsets any flags that were previously set by
// [YoutubeIncludeDashManifest].
func (c *Command) UnsetYoutubeIncludeDashManifest() *Command {
	c.removeFlagByID("youtube_include_dash_manifest")
	return c
}

// YoutubeSkipDashManifest sets the "youtube-skip-dash-manifest" flag (no description specified).
//
//   - See also [UnsetYoutubeSkipDashManifest], for unsetting the flag.
//   - YoutubeSkipDashManifest maps to cli flags: --youtube-skip-dash-manifest/--no-youtube-include-dash-manifest (hidden).
func (c *Command) YoutubeSkipDashManifest() *Command {
	c.addFlag(&Flag{
		ID:   "youtube_include_dash_manifest",
		Flag: "--youtube-skip-dash-manifest",
		Args: nil,
	})
	return c
}

// UnsetYoutubeSkipDashManifest unsets any flags that were previously set by
// [YoutubeSkipDashManifest].
func (c *Command) UnsetYoutubeSkipDashManifest() *Command {
	c.removeFlagByID("youtube_include_dash_manifest")
	return c
}

// YoutubeIncludeHlsManifest sets the "youtube-include-hls-manifest" flag (no description specified).
//
//   - See also [UnsetYoutubeIncludeHlsManifest], for unsetting the flag.
//   - YoutubeIncludeHlsManifest maps to cli flags: --youtube-include-hls-manifest/--no-youtube-skip-hls-manifest (hidden).
func (c *Command) YoutubeIncludeHlsManifest() *Command {
	c.addFlag(&Flag{
		ID:   "youtube_include_hls_manifest",
		Flag: "--youtube-include-hls-manifest",
		Args: nil,
	})
	return c
}

// UnsetYoutubeIncludeHlsManifest unsets any flags that were previously set by
// [YoutubeIncludeHlsManifest].
func (c *Command) UnsetYoutubeIncludeHlsManifest() *Command {
	c.removeFlagByID("youtube_include_hls_manifest")
	return c
}

// YoutubeSkipHlsManifest sets the "youtube-skip-hls-manifest" flag (no description specified).
//
//   - See also [UnsetYoutubeSkipHlsManifest], for unsetting the flag.
//   - YoutubeSkipHlsManifest maps to cli flags: --youtube-skip-hls-manifest/--no-youtube-include-hls-manifest (hidden).
func (c *Command) YoutubeSkipHlsManifest() *Command {
	c.addFlag(&Flag{
		ID:   "youtube_include_hls_manifest",
		Flag: "--youtube-skip-hls-manifest",
		Args: nil,
	})
	return c
}

// UnsetYoutubeSkipHlsManifest unsets any flags that were previously set by
// [YoutubeSkipHlsManifest].
func (c *Command) UnsetYoutubeSkipHlsManifest() *Command {
	c.removeFlagByID("youtube_include_hls_manifest")
	return c
}

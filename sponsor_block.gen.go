// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by cmd/codegen. DO NOT EDIT.
//
// SponsorBlock Option Group -- Description:
//   Make chapter entries for, or remove various segments (sponsor, introductions,
//   etc.) from downloaded YouTube videos using the SponsorBlock API
//   (https://sponsor.ajay.app)

package ytdlp

// SponsorBlock categories to create chapters for, separated by commas. Available
// categories are sponsor, intro, outro, selfpromo, preview, filler, interaction,
// music_offtopic, poi_highlight, chapter, all and default (=all). You can prefix
// the category with a "-" to exclude it. See [1] for description of the
// categories. E.g. --sponsorblock-mark all,-preview [1]
// https://wiki.sponsor.ajay.app/w/Segment_Categories
//
//   - See also [UnsetSponsorblockMark], for unsetting the flag.
//   - SponsorblockMark maps to cli flags: --sponsorblock-mark=CATS.
func (c *Command) SponsorblockMark(cats string) *Command {
	c.addFlag(&Flag{
		ID:   "sponsorblock_mark",
		Flag: "--sponsorblock-mark",
		Args: []string{cats},
	})
	return c
}

// UnsetSponsorblockMark unsets any flags that were previously set by
// [SponsorblockMark].
func (c *Command) UnsetSponsorblockMark() *Command {
	c.removeFlagByID("sponsorblock_mark")
	return c
}

// SponsorBlock categories to be removed from the video file, separated by commas.
// If a category is present in both mark and remove, remove takes precedence. The
// syntax and available categories are the same as for --sponsorblock-mark except
// that "default" refers to "all,-filler" and poi_highlight, chapter are not
// available
//
//   - See also [UnsetSponsorblockRemove], for unsetting the flag.
//   - SponsorblockRemove maps to cli flags: --sponsorblock-remove=CATS.
func (c *Command) SponsorblockRemove(cats string) *Command {
	c.addFlag(&Flag{
		ID:   "sponsorblock_remove",
		Flag: "--sponsorblock-remove",
		Args: []string{cats},
	})
	return c
}

// UnsetSponsorblockRemove unsets any flags that were previously set by
// [SponsorblockRemove].
func (c *Command) UnsetSponsorblockRemove() *Command {
	c.removeFlagByID("sponsorblock_remove")
	return c
}

// An output template for the title of the SponsorBlock chapters created by
// --sponsorblock-mark. The only available fields are start_time, end_time,
// category, categories, name, category_names. Defaults to "[SponsorBlock]:
// %(category_names)l"
//
//   - See also [UnsetSponsorblockChapterTitle], for unsetting the flag.
//   - SponsorblockChapterTitle maps to cli flags: --sponsorblock-chapter-title=TEMPLATE.
func (c *Command) SponsorblockChapterTitle(template string) *Command {
	c.addFlag(&Flag{
		ID:   "sponsorblock_chapter_title",
		Flag: "--sponsorblock-chapter-title",
		Args: []string{template},
	})
	return c
}

// UnsetSponsorblockChapterTitle unsets any flags that were previously set by
// [SponsorblockChapterTitle].
func (c *Command) UnsetSponsorblockChapterTitle() *Command {
	c.removeFlagByID("sponsorblock_chapter_title")
	return c
}

// Disable both --sponsorblock-mark and --sponsorblock-remove
//
//   - See also [UnsetNoSponsorblock], for unsetting the flag.
//   - NoSponsorblock maps to cli flags: --no-sponsorblock.
func (c *Command) NoSponsorblock() *Command {
	c.addFlag(&Flag{
		ID:   "no_sponsorblock",
		Flag: "--no-sponsorblock",
		Args: nil,
	})
	return c
}

// UnsetNoSponsorblock unsets any flags that were previously set by
// [NoSponsorblock].
func (c *Command) UnsetNoSponsorblock() *Command {
	c.removeFlagByID("no_sponsorblock")
	return c
}

// SponsorBlock API location, defaults to https://sponsor.ajay.app
//
//   - See also [UnsetSponsorblockApi], for unsetting the flag.
//   - SponsorblockApi maps to cli flags: --sponsorblock-api=URL.
func (c *Command) SponsorblockApi(url string) *Command {
	c.addFlag(&Flag{
		ID:   "sponsorblock_api",
		Flag: "--sponsorblock-api",
		Args: []string{url},
	})
	return c
}

// UnsetSponsorblockApi unsets any flags that were previously set by
// [SponsorblockApi].
func (c *Command) UnsetSponsorblockApi() *Command {
	c.removeFlagByID("sponsorblock_api")
	return c
}

// Sponskrub sets the "sponskrub" flag (no description specified).
//
//   - See also [UnsetSponskrub], for unsetting the flag.
//   - Sponskrub maps to cli flags: --sponskrub (hidden).
func (c *Command) Sponskrub() *Command {
	c.addFlag(&Flag{
		ID:   "sponskrub",
		Flag: "--sponskrub",
		Args: nil,
	})
	return c
}

// UnsetSponskrub unsets any flags that were previously set by
// [Sponskrub].
func (c *Command) UnsetSponskrub() *Command {
	c.removeFlagByID("sponskrub")
	return c
}

// NoSponskrub sets the "no-sponskrub" flag (no description specified).
//
//   - See also [UnsetNoSponskrub], for unsetting the flag.
//   - NoSponskrub maps to cli flags: --no-sponskrub (hidden).
func (c *Command) NoSponskrub() *Command {
	c.addFlag(&Flag{
		ID:   "sponskrub",
		Flag: "--no-sponskrub",
		Args: nil,
	})
	return c
}

// UnsetNoSponskrub unsets any flags that were previously set by
// [NoSponskrub].
func (c *Command) UnsetNoSponskrub() *Command {
	c.removeFlagByID("sponskrub")
	return c
}

// SponskrubCut sets the "sponskrub-cut" flag (no description specified).
//
//   - See also [UnsetSponskrubCut], for unsetting the flag.
//   - SponskrubCut maps to cli flags: --sponskrub-cut (hidden).
func (c *Command) SponskrubCut() *Command {
	c.addFlag(&Flag{
		ID:   "sponskrub_cut",
		Flag: "--sponskrub-cut",
		Args: nil,
	})
	return c
}

// UnsetSponskrubCut unsets any flags that were previously set by
// [SponskrubCut].
func (c *Command) UnsetSponskrubCut() *Command {
	c.removeFlagByID("sponskrub_cut")
	return c
}

// NoSponskrubCut sets the "no-sponskrub-cut" flag (no description specified).
//
//   - See also [UnsetNoSponskrubCut], for unsetting the flag.
//   - NoSponskrubCut maps to cli flags: --no-sponskrub-cut (hidden).
func (c *Command) NoSponskrubCut() *Command {
	c.addFlag(&Flag{
		ID:   "sponskrub_cut",
		Flag: "--no-sponskrub-cut",
		Args: nil,
	})
	return c
}

// UnsetNoSponskrubCut unsets any flags that were previously set by
// [NoSponskrubCut].
func (c *Command) UnsetNoSponskrubCut() *Command {
	c.removeFlagByID("sponskrub_cut")
	return c
}

// SponskrubForce sets the "sponskrub-force" flag (no description specified).
//
//   - See also [UnsetSponskrubForce], for unsetting the flag.
//   - SponskrubForce maps to cli flags: --sponskrub-force (hidden).
func (c *Command) SponskrubForce() *Command {
	c.addFlag(&Flag{
		ID:   "sponskrub_force",
		Flag: "--sponskrub-force",
		Args: nil,
	})
	return c
}

// UnsetSponskrubForce unsets any flags that were previously set by
// [SponskrubForce].
func (c *Command) UnsetSponskrubForce() *Command {
	c.removeFlagByID("sponskrub_force")
	return c
}

// NoSponskrubForce sets the "no-sponskrub-force" flag (no description specified).
//
//   - See also [UnsetNoSponskrubForce], for unsetting the flag.
//   - NoSponskrubForce maps to cli flags: --no-sponskrub-force (hidden).
func (c *Command) NoSponskrubForce() *Command {
	c.addFlag(&Flag{
		ID:   "sponskrub_force",
		Flag: "--no-sponskrub-force",
		Args: nil,
	})
	return c
}

// UnsetNoSponskrubForce unsets any flags that were previously set by
// [NoSponskrubForce].
func (c *Command) UnsetNoSponskrubForce() *Command {
	c.removeFlagByID("sponskrub_force")
	return c
}

// - See also [UnsetSponskrubLocation], for unsetting the flag.
// - SponskrubLocation maps to cli flags: --sponskrub-location=PATH (hidden).
func (c *Command) SponskrubLocation(path string) *Command {
	c.addFlag(&Flag{
		ID:   "sponskrub_path",
		Flag: "--sponskrub-location",
		Args: []string{path},
	})
	return c
}

// UnsetSponskrubLocation unsets any flags that were previously set by
// [SponskrubLocation].
func (c *Command) UnsetSponskrubLocation() *Command {
	c.removeFlagByID("sponskrub_path")
	return c
}

// - See also [UnsetSponskrubArgs], for unsetting the flag.
// - SponskrubArgs maps to cli flags: --sponskrub-args=ARGS (hidden).
func (c *Command) SponskrubArgs(args string) *Command {
	c.addFlag(&Flag{
		ID:   "sponskrub_args",
		Flag: "--sponskrub-args",
		Args: []string{args},
	})
	return c
}

// UnsetSponskrubArgs unsets any flags that were previously set by
// [SponskrubArgs].
func (c *Command) UnsetSponskrubArgs() *Command {
	c.removeFlagByID("sponskrub_args")
	return c
}

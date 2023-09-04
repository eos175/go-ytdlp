// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by cmd/codegen. DO NOT EDIT.
//
// Verbosity Simulation Option Group

package ytdlp

// Activate quiet mode. If used with --verbose, print the log to stderr
//
//   - See [UnsetQuiet], for unsetting the flag.
//   - Quiet maps to cli flags: -q/--quiet.
func (c *Command) Quiet() *Command {
	c.addFlag(&Flag{
		ID:   "quiet",
		Flag: "--quiet",
		Args: nil,
	})
	return c
}

// UnsetQuiet unsets any flags that were previously set by
// [Quiet].
func (c *Command) UnsetQuiet() *Command {
	c.removeFlagByID("quiet")
	return c
}

// Deactivate quiet mode. (Default)
//
//   - See [UnsetQuiet], for unsetting the flag.
//   - NoQuiet maps to cli flags: --no-quiet.
func (c *Command) NoQuiet() *Command {
	c.addFlag(&Flag{
		ID:   "quiet",
		Flag: "--no-quiet",
		Args: nil,
	})
	return c
}

// Ignore warnings
//
//   - See [UnsetWarnings], for unsetting the flag.
//   - NoWarnings maps to cli flags: --no-warnings.
func (c *Command) NoWarnings() *Command {
	c.addFlag(&Flag{
		ID:   "no_warnings",
		Flag: "--no-warnings",
		Args: nil,
	})
	return c
}

// Do not download the video and do not write anything to disk
//
//   - See [UnsetSimulate], for unsetting the flag.
//   - Simulate maps to cli flags: -s/--simulate.
func (c *Command) Simulate() *Command {
	c.addFlag(&Flag{
		ID:   "simulate",
		Flag: "--simulate",
		Args: nil,
	})
	return c
}

// UnsetSimulate unsets any flags that were previously set by
// [Simulate].
func (c *Command) UnsetSimulate() *Command {
	c.removeFlagByID("simulate")
	return c
}

// Download the video even if printing/listing options are used
//
//   - See [UnsetSimulate], for unsetting the flag.
//   - NoSimulate maps to cli flags: --no-simulate.
func (c *Command) NoSimulate() *Command {
	c.addFlag(&Flag{
		ID:   "simulate",
		Flag: "--no-simulate",
		Args: nil,
	})
	return c
}

// Ignore "No video formats" error. Useful for extracting metadata even if the
// videos are not actually available for download (experimental)
//
//   - See [UnsetIgnoreNoFormatsError], for unsetting the flag.
//   - IgnoreNoFormatsError maps to cli flags: --ignore-no-formats-error.
func (c *Command) IgnoreNoFormatsError() *Command {
	c.addFlag(&Flag{
		ID:   "ignore_no_formats_error",
		Flag: "--ignore-no-formats-error",
		Args: nil,
	})
	return c
}

// UnsetIgnoreNoFormatsError unsets any flags that were previously set by
// [IgnoreNoFormatsError].
func (c *Command) UnsetIgnoreNoFormatsError() *Command {
	c.removeFlagByID("ignore_no_formats_error")
	return c
}

// Throw error when no downloadable video formats are found (default)
//
//   - See [UnsetIgnoreNoFormatsError], for unsetting the flag.
//   - NoIgnoreNoFormatsError maps to cli flags: --no-ignore-no-formats-error.
func (c *Command) NoIgnoreNoFormatsError() *Command {
	c.addFlag(&Flag{
		ID:   "ignore_no_formats_error",
		Flag: "--no-ignore-no-formats-error",
		Args: nil,
	})
	return c
}

// Do not download the video but write all related files
//
//   - See [UnsetSkipDownload], for unsetting the flag.
//   - SkipDownload maps to cli flags: --skip-download/--no-download.
func (c *Command) SkipDownload() *Command {
	c.addFlag(&Flag{
		ID:   "skip_download",
		Flag: "--skip-download",
		Args: nil,
	})
	return c
}

// UnsetSkipDownload unsets any flags that were previously set by
// [SkipDownload].
func (c *Command) UnsetSkipDownload() *Command {
	c.removeFlagByID("skip_download")
	return c
}

// Field name or output template to print to screen, optionally prefixed with when
// to print it, separated by a ":". Supported values of "WHEN" are the same as that
// of --use-postprocessor (default: video). Implies --quiet. Implies --simulate
// unless --no-simulate or later stages of WHEN are used. This option can be used
// multiple times
//
//   - See [UnsetPrint], for unsetting the flag.
//   - Print maps to cli flags: -O/--print=[WHEN:]TEMPLATE.
func (c *Command) Print(template string) *Command {
	c.addFlag(&Flag{
		ID:   "forceprint",
		Flag: "--print",
		Args: []string{template},
	})
	return c
}

// UnsetPrint unsets any flags that were previously set by
// [Print].
func (c *Command) UnsetPrint() *Command {
	c.removeFlagByID("forceprint")
	return c
}

// Append given template to the file. The values of WHEN and TEMPLATE are same as
// that of --print. FILE uses the same syntax as the output template. This option
// can be used multiple times
//
//   - See [UnsetPrintToFile], for unsetting the flag.
//   - PrintToFile maps to cli flags: --print-to-file=[WHEN:]TEMPLATE FILE.
func (c *Command) PrintToFile(template, file string) *Command {
	c.addFlag(&Flag{
		ID:   "print_to_file",
		Flag: "--print-to-file",
		Args: []string{template, file},
	})
	return c
}

// UnsetPrintToFile unsets any flags that were previously set by
// [PrintToFile].
func (c *Command) UnsetPrintToFile() *Command {
	c.removeFlagByID("print_to_file")
	return c
}

// GetUrl sets the "get-url" flag (no description specified).
//
//   - See [UnsetGetUrl], for unsetting the flag.
//   - GetUrl maps to cli flags: -g/--get-url (hidden).
func (c *Command) GetUrl() *Command {
	c.addFlag(&Flag{
		ID:   "geturl",
		Flag: "--get-url",
		Args: nil,
	})
	return c
}

// UnsetGetUrl unsets any flags that were previously set by
// [GetUrl].
func (c *Command) UnsetGetUrl() *Command {
	c.removeFlagByID("geturl")
	return c
}

// GetTitle sets the "get-title" flag (no description specified).
//
//   - See [UnsetGetTitle], for unsetting the flag.
//   - GetTitle maps to cli flags: -e/--get-title (hidden).
func (c *Command) GetTitle() *Command {
	c.addFlag(&Flag{
		ID:   "gettitle",
		Flag: "--get-title",
		Args: nil,
	})
	return c
}

// UnsetGetTitle unsets any flags that were previously set by
// [GetTitle].
func (c *Command) UnsetGetTitle() *Command {
	c.removeFlagByID("gettitle")
	return c
}

// GetId sets the "get-id" flag (no description specified).
//
//   - See [UnsetGetId], for unsetting the flag.
//   - GetId maps to cli flags: --get-id (hidden).
func (c *Command) GetId() *Command {
	c.addFlag(&Flag{
		ID:   "getid",
		Flag: "--get-id",
		Args: nil,
	})
	return c
}

// UnsetGetId unsets any flags that were previously set by
// [GetId].
func (c *Command) UnsetGetId() *Command {
	c.removeFlagByID("getid")
	return c
}

// GetThumbnail sets the "get-thumbnail" flag (no description specified).
//
//   - See [UnsetGetThumbnail], for unsetting the flag.
//   - GetThumbnail maps to cli flags: --get-thumbnail (hidden).
func (c *Command) GetThumbnail() *Command {
	c.addFlag(&Flag{
		ID:   "getthumbnail",
		Flag: "--get-thumbnail",
		Args: nil,
	})
	return c
}

// UnsetGetThumbnail unsets any flags that were previously set by
// [GetThumbnail].
func (c *Command) UnsetGetThumbnail() *Command {
	c.removeFlagByID("getthumbnail")
	return c
}

// GetDescription sets the "get-description" flag (no description specified).
//
//   - See [UnsetGetDescription], for unsetting the flag.
//   - GetDescription maps to cli flags: --get-description (hidden).
func (c *Command) GetDescription() *Command {
	c.addFlag(&Flag{
		ID:   "getdescription",
		Flag: "--get-description",
		Args: nil,
	})
	return c
}

// UnsetGetDescription unsets any flags that were previously set by
// [GetDescription].
func (c *Command) UnsetGetDescription() *Command {
	c.removeFlagByID("getdescription")
	return c
}

// GetDuration sets the "get-duration" flag (no description specified).
//
//   - See [UnsetGetDuration], for unsetting the flag.
//   - GetDuration maps to cli flags: --get-duration (hidden).
func (c *Command) GetDuration() *Command {
	c.addFlag(&Flag{
		ID:   "getduration",
		Flag: "--get-duration",
		Args: nil,
	})
	return c
}

// UnsetGetDuration unsets any flags that were previously set by
// [GetDuration].
func (c *Command) UnsetGetDuration() *Command {
	c.removeFlagByID("getduration")
	return c
}

// GetFilename sets the "get-filename" flag (no description specified).
//
//   - See [UnsetGetFilename], for unsetting the flag.
//   - GetFilename maps to cli flags: --get-filename (hidden).
func (c *Command) GetFilename() *Command {
	c.addFlag(&Flag{
		ID:   "getfilename",
		Flag: "--get-filename",
		Args: nil,
	})
	return c
}

// UnsetGetFilename unsets any flags that were previously set by
// [GetFilename].
func (c *Command) UnsetGetFilename() *Command {
	c.removeFlagByID("getfilename")
	return c
}

// GetFormat sets the "get-format" flag (no description specified).
//
//   - See [UnsetGetFormat], for unsetting the flag.
//   - GetFormat maps to cli flags: --get-format (hidden).
func (c *Command) GetFormat() *Command {
	c.addFlag(&Flag{
		ID:   "getformat",
		Flag: "--get-format",
		Args: nil,
	})
	return c
}

// UnsetGetFormat unsets any flags that were previously set by
// [GetFormat].
func (c *Command) UnsetGetFormat() *Command {
	c.removeFlagByID("getformat")
	return c
}

// Quiet, but print JSON information for each video. Simulate unless --no-simulate
// is used. See "OUTPUT TEMPLATE" for a description of available keys
//
//   - See [UnsetDumpJson], for unsetting the flag.
//   - DumpJson maps to cli flags: -j/--dump-json.
func (c *Command) DumpJson() *Command {
	c.addFlag(&Flag{
		ID:   "dumpjson",
		Flag: "--dump-json",
		Args: nil,
	})
	return c
}

// UnsetDumpJson unsets any flags that were previously set by
// [DumpJson].
func (c *Command) UnsetDumpJson() *Command {
	c.removeFlagByID("dumpjson")
	return c
}

// Quiet, but print JSON information for each url or infojson passed. Simulate
// unless --no-simulate is used. If the URL refers to a playlist, the whole
// playlist information is dumped in a single line
//
//   - See [UnsetDumpSingleJson], for unsetting the flag.
//   - DumpSingleJson maps to cli flags: -J/--dump-single-json.
func (c *Command) DumpSingleJson() *Command {
	c.addFlag(&Flag{
		ID:   "dump_single_json",
		Flag: "--dump-single-json",
		Args: nil,
	})
	return c
}

// UnsetDumpSingleJson unsets any flags that were previously set by
// [DumpSingleJson].
func (c *Command) UnsetDumpSingleJson() *Command {
	c.removeFlagByID("dump_single_json")
	return c
}

// PrintJson sets the "print-json" flag (no description specified).
//
//   - See [UnsetPrintJson], for unsetting the flag.
//   - PrintJson maps to cli flags: --print-json (hidden).
func (c *Command) PrintJson() *Command {
	c.addFlag(&Flag{
		ID:   "print_json",
		Flag: "--print-json",
		Args: nil,
	})
	return c
}

// UnsetPrintJson unsets any flags that were previously set by
// [PrintJson].
func (c *Command) UnsetPrintJson() *Command {
	c.removeFlagByID("print_json")
	return c
}

// Force download archive entries to be written as far as no errors occur, even if
// -s or another simulation option is used
//
//   - See [UnsetForceWriteArchive], for unsetting the flag.
//   - ForceWriteArchive maps to cli flags: --force-write-archive/--force-write-download-archive/--force-download-archive.
func (c *Command) ForceWriteArchive() *Command {
	c.addFlag(&Flag{
		ID:   "force_write_download_archive",
		Flag: "--force-write-archive",
		Args: nil,
	})
	return c
}

// UnsetForceWriteArchive unsets any flags that were previously set by
// [ForceWriteArchive].
func (c *Command) UnsetForceWriteArchive() *Command {
	c.removeFlagByID("force_write_download_archive")
	return c
}

// Output progress bar as new lines
//
//   - See [UnsetNewline], for unsetting the flag.
//   - Newline maps to cli flags: --newline.
func (c *Command) Newline() *Command {
	c.addFlag(&Flag{
		ID:   "progress_with_newline",
		Flag: "--newline",
		Args: nil,
	})
	return c
}

// UnsetNewline unsets any flags that were previously set by
// [Newline].
func (c *Command) UnsetNewline() *Command {
	c.removeFlagByID("progress_with_newline")
	return c
}

// Do not print progress bar
//
//   - See [UnsetProgress], for unsetting the flag.
//   - NoProgress maps to cli flags: --no-progress.
func (c *Command) NoProgress() *Command {
	c.addFlag(&Flag{
		ID:   "noprogress",
		Flag: "--no-progress",
		Args: nil,
	})
	return c
}

// Show progress bar, even if in quiet mode
//
//   - See [UnsetProgress], for unsetting the flag.
//   - Progress maps to cli flags: --progress.
func (c *Command) Progress() *Command {
	c.addFlag(&Flag{
		ID:   "noprogress",
		Flag: "--progress",
		Args: nil,
	})
	return c
}

// UnsetProgress unsets any flags that were previously set by
// [Progress].
func (c *Command) UnsetProgress() *Command {
	c.removeFlagByID("noprogress")
	return c
}

// Display progress in console titlebar
//
//   - See [UnsetConsoleTitle], for unsetting the flag.
//   - ConsoleTitle maps to cli flags: --console-title.
func (c *Command) ConsoleTitle() *Command {
	c.addFlag(&Flag{
		ID:   "consoletitle",
		Flag: "--console-title",
		Args: nil,
	})
	return c
}

// UnsetConsoleTitle unsets any flags that were previously set by
// [ConsoleTitle].
func (c *Command) UnsetConsoleTitle() *Command {
	c.removeFlagByID("consoletitle")
	return c
}

// Template for progress outputs, optionally prefixed with one of "download:"
// (default), "download-title:" (the console title), "postprocess:",  or
// "postprocess-title:". The video's fields are accessible under the "info" key and
// the progress attributes are accessible under "progress" key. E.g.
// --console-title --progress-template
// "download-title:%(info.id)s-%(progress.eta)s"
//
//   - See [UnsetProgressTemplate], for unsetting the flag.
//   - ProgressTemplate maps to cli flags: --progress-template=[TYPES:]TEMPLATE.
func (c *Command) ProgressTemplate(template string) *Command {
	c.addFlag(&Flag{
		ID:   "progress_template",
		Flag: "--progress-template",
		Args: []string{template},
	})
	return c
}

// UnsetProgressTemplate unsets any flags that were previously set by
// [ProgressTemplate].
func (c *Command) UnsetProgressTemplate() *Command {
	c.removeFlagByID("progress_template")
	return c
}

// Print various debugging information
//
//   - See [UnsetVerbose], for unsetting the flag.
//   - Verbose maps to cli flags: -v/--verbose.
func (c *Command) Verbose() *Command {
	c.addFlag(&Flag{
		ID:   "verbose",
		Flag: "--verbose",
		Args: nil,
	})
	return c
}

// UnsetVerbose unsets any flags that were previously set by
// [Verbose].
func (c *Command) UnsetVerbose() *Command {
	c.removeFlagByID("verbose")
	return c
}

// Print downloaded pages encoded using base64 to debug problems (very verbose)
//
//   - See [UnsetDumpPages], for unsetting the flag.
//   - DumpPages maps to cli flags: --dump-pages/--dump-intermediate-pages.
func (c *Command) DumpPages() *Command {
	c.addFlag(&Flag{
		ID:   "dump_intermediate_pages",
		Flag: "--dump-pages",
		Args: nil,
	})
	return c
}

// UnsetDumpPages unsets any flags that were previously set by
// [DumpPages].
func (c *Command) UnsetDumpPages() *Command {
	c.removeFlagByID("dump_intermediate_pages")
	return c
}

// Write downloaded intermediary pages to files in the current directory to debug
// problems
//
//   - See [UnsetWritePages], for unsetting the flag.
//   - WritePages maps to cli flags: --write-pages.
func (c *Command) WritePages() *Command {
	c.addFlag(&Flag{
		ID:   "write_pages",
		Flag: "--write-pages",
		Args: nil,
	})
	return c
}

// UnsetWritePages unsets any flags that were previously set by
// [WritePages].
func (c *Command) UnsetWritePages() *Command {
	c.removeFlagByID("write_pages")
	return c
}

// LoadPages sets the "load-pages" flag (no description specified).
//
//   - See [UnsetLoadPages], for unsetting the flag.
//   - LoadPages maps to cli flags: --load-pages (hidden).
func (c *Command) LoadPages() *Command {
	c.addFlag(&Flag{
		ID:   "load_pages",
		Flag: "--load-pages",
		Args: nil,
	})
	return c
}

// UnsetLoadPages unsets any flags that were previously set by
// [LoadPages].
func (c *Command) UnsetLoadPages() *Command {
	c.removeFlagByID("load_pages")
	return c
}

// YoutubePrintSigCode sets the "youtube-print-sig-code" flag (no description specified).
//
//   - See [UnsetYoutubePrintSigCode], for unsetting the flag.
//   - YoutubePrintSigCode maps to cli flags: --youtube-print-sig-code (hidden).
func (c *Command) YoutubePrintSigCode() *Command {
	c.addFlag(&Flag{
		ID:   "youtube_print_sig_code",
		Flag: "--youtube-print-sig-code",
		Args: nil,
	})
	return c
}

// UnsetYoutubePrintSigCode unsets any flags that were previously set by
// [YoutubePrintSigCode].
func (c *Command) UnsetYoutubePrintSigCode() *Command {
	c.removeFlagByID("youtube_print_sig_code")
	return c
}

// Display sent and read HTTP traffic
//
//   - See [UnsetPrintTraffic], for unsetting the flag.
//   - PrintTraffic maps to cli flags: --print-traffic/--dump-headers.
func (c *Command) PrintTraffic() *Command {
	c.addFlag(&Flag{
		ID:   "debug_printtraffic",
		Flag: "--print-traffic",
		Args: nil,
	})
	return c
}

// UnsetPrintTraffic unsets any flags that were previously set by
// [PrintTraffic].
func (c *Command) UnsetPrintTraffic() *Command {
	c.removeFlagByID("debug_printtraffic")
	return c
}

// CallHome sets the "call-home" flag (no description specified).
//
//   - See [UnsetCallHome], for unsetting the flag.
//   - CallHome maps to cli flags: -C/--call-home (hidden).
func (c *Command) CallHome() *Command {
	c.addFlag(&Flag{
		ID:   "call_home",
		Flag: "--call-home",
		Args: nil,
	})
	return c
}

// UnsetCallHome unsets any flags that were previously set by
// [CallHome].
func (c *Command) UnsetCallHome() *Command {
	c.removeFlagByID("call_home")
	return c
}

// NoCallHome sets the "no-call-home" flag (no description specified).
//
//   - See [UnsetCallHome], for unsetting the flag.
//   - NoCallHome maps to cli flags: --no-call-home (hidden).
func (c *Command) NoCallHome() *Command {
	c.addFlag(&Flag{
		ID:   "call_home",
		Flag: "--no-call-home",
		Args: nil,
	})
	return c
}

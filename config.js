/*
	Config file for a basic IRC bridge plugin written in GO
	Binary will be spawned using child_process.spawn and get input from stdin and text from stdout will be ran as
	commands in the factorio server.
*/
var os = require("os");

var binary = "build/clusterioIRCbridge_windows";
if(os.platform == "linux") binary = "build/clusterioIRCbridge_linux";
if(os.platform == "darwin") binary = "build/clusterioIRCbridge_darwin";

module.exports = {
	// Name of package. For display somewhere I guess.
	name: "clusterioIRCbridge",
	version: "1.0.0",
	// Binary entrypoint for plugin. Don't let it crash. Stdout is sent to game as server chat (run /c commands from here for interface)
	// Make sure its cross platform somehow.
	binary: binary,
	description: "Clusterio plugin to put #factorio on espernet in your game chat",
	// We'll send everything in this file to your stdin. Beware.
	scriptOutputFileSubscription: false,
}
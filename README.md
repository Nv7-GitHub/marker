# Marker

Welcome to marker, a finder utility for remembering where your files are!
*WARNING\: This is only meant for Mac. I may create marker for Windows and Linux when I get a Windows & Linux PC*

## Installation
You can get a prebuilt package in the [releases](https://github.com/Nv7-GitHub/marker/releases/) section.
To install, download the pkg and double-click it. This will bring up an installation window.
![Window](https://i.imgur.com/X70QwNT.png)
Press continue, and then press Install.
![Install button](https://i.imgur.com/LIqs0Ah.png)
This will ask you to enter in an administrator password, which is required. Please enter an administrator password.
This will lead to 2 notifications (below). Press Install on both.
![Notifications](https://i.imgur.com/OKTSkP1.png)
Now it is installed!

## Usage
Marker allows you to find your files. To add a marker to a file or folder, open Finder and select a folder you want to put a marker on.
![Folder selected](https://i.imgur.com/yTrQNW0.png)
Right click on the folder, and hover over "Quick Actions". If you have installed Marker properly, there should be two options: "New Marker", and "View Markers"
![Options](https://i.imgur.com/KfYx7Qz.png)
Press NewMarker to make a new marker on that folder. Press ViewMarkers to view, edit, and remove markers.

### NewMarker
NewMarker is pretty self explanatory. Just put a title and a description.

### ViewMarkers
ViewMarkers allows you to view your markers. Click on a marker to view it. Press on the Folder icon to reveal in finder, and and the trash icon to delete the marker. Press the back button to return to the view of all the markers.

### Themes
There is both a light theme and a dark theme in marker. It is by default in the dark theme. You can change to the light theme by going to the menu, and pressing theme. In that menu, there are two buttons, Light theme and dark theme

## Deleting Marker
To delete Marker, delete the following items. Remember to replace ```<your user>``` with your username
- In /Users/Shared delete the "Marker" folder.
- In /Users/```<your user>```/Library/Services, delete "NewMarker" and "ViewMarkers"

## Compiling from source
To compile from source, download this repository with ```git clone```, or just download as zip.
Next, you need to compile the code. To do this, navigate inside the folder with the code and run the following commands:
```bash
go get fyne.io/fyne
go build -o marker main.go new.go find.go
```
Once you have done this, you need to move files. Make a folder in the ```Shared``` user called Marker and copy the compiled output, a file called "marker" into the folder you just made. Finally, double-click on both the workflows in the workflows folder and press install in the notification.

## Thank you for using Marker!

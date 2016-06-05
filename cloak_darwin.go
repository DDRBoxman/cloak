package cloak

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework AppKit
#import <AppKit/AppKit.h>

void
hide(void) {
	[NSCursor setHiddenUntilMouseMoves:YES];
}

*/
import "C"

func Hide() {
	C.hide()
}

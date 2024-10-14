////////////////////////////////////////////////////////////////////////

package tools

import (
    "fmt"
    cv "gocv.io/x/gocv"
)

////////////////////////////////////////////////////////////////////////


func OpenCV() {
    fmt.Println("7-3. OpenCV Testing ...")

    webcam, err := cv.VideoCaptureDevice(0)
    if err != nil {
        fmt.Printf("Error opening capture device: %v", err)
        return
    }
    defer webcam.Close()

    // Create a new window to display the video
    window := cv.NewWindow("Real-time camera feed")
    defer window.Close()

    // Matrix to hold image frame data
    img := cv.NewMat()
    defer img.Close()

    for {
        if ok := webcam.Read(&img); !ok {
            fmt.Printf("Error reading webcam\n")
            return
        }
        if img.Empty() {
            return
        }

        // Display the image in the window
        window.IMShow(img)

        // Wait for 1ms and exit the loop if `ESC` key is pressed
        if window.WaitKey(1) == 27 {
            break
        }
    }
}



package editor

func (tf *TextField) trimCursor() {
    c := tf.cursors[0]
    c.pos.Y = max(0, min(c.pos.Y, len(tf.text)))
    c.pos.X = max(0, min(c.pos.X, len(tf.text[c.pos.Y])))
}

func (tf *TextField) MoveCursor(x, y int) {
    c := tf.cursors[0]
    c.Move(x, y)
    c.pos.Y = max(0, min(c.pos.Y, len(tf.text)))
}

func (tf *TextField) InsertLine(i int) {
    tf.text = append(tf.text[:i+1], append([]string{""}, tf.text[i+1:]...)...)
}

func (tf *TextField) HandleHandleKeyPress(key string, ctrl bool, shift bool, alt bool) {
    switch key {
        case "Enter":
            tf.InsertLine(tf.cursors[0].pos.Y)
            tf.MoveCursor(0, 1)
            return;
        case "ArrowLeft":
            if tf.cursors[0].pos.X > 0 {
                tf.MoveCursor(-1, 0)
            }
            return
        case "ArrowRight":
            if tf.cursors[0].pos.X < len(tf.text[tf.cursors[0].pos.Y]) {
                tf.MoveCursor(1, 0)
            }
            return
        case "ArrowUp":
            if tf.cursors[0].pos.Y > 0 {
                tf.MoveCursor(0, -1)
            }
            return
        case "ArrowDown":
            if tf.cursors[0].pos.Y < len(tf.text) - 1 {
                tf.MoveCursor(0, 1)
            }
            return
    }
    tf.trimCursor()
    switch key {
        case "Backspace":
            if tf.cursors[0].pos.X > 0 {
                tf.text[tf.cursors[0].pos.Y] = tf.text[tf.cursors[0].pos.Y][:tf.cursors[0].pos.X-1] +
                                               tf.text[tf.cursors[0].pos.Y][tf.cursors[0].pos.X:]
                tf.MoveCursor(-1, 0)
            }
            return
        case "Delete":
            if tf.cursors[0].pos.X < len(tf.text[tf.cursors[0].pos.Y]) {
                tf.text[tf.cursors[0].pos.Y] = tf.text[tf.cursors[0].pos.Y][:tf.cursors[0].pos.X] +
                                               tf.text[tf.cursors[0].pos.Y][tf.cursors[0].pos.X+1:]
            }
            return
    }
    if len(key) != 1 {
        return
    }
    tf.text[tf.cursors[0].pos.Y] = tf.text[tf.cursors[0].pos.Y][:tf.cursors[0].pos.X] + key + tf.text[tf.cursors[0].pos.Y][tf.cursors[0].pos.X:]
    tf.MoveCursor(1, 0)
}

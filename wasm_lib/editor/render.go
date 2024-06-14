package editor

import (
    "github.com/gomarkdown/markdown"
)

func (tf *TextField) Render() []string {
    res := make([]string, 0)

    for i, md_line := range tf.text {
        c := tf.cursors[0]
        if c.Pos().Y != i {
            res = append(res, string(markdown.ToHTML([]byte(md_line), nil, nil)))
            continue
        }
        x := max(0, min(c.Pos().X, len(md_line)))
        bf := md_line[:x]
        af := md_line[x:]
        middle := "<div class='cursor'></div>"
        res = append(res, bf + middle + af)
    }
    
    return res
}

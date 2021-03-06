/*
 *
 * basic2d.go - A library for 2D primitives using image/draw
 *   Copyright Brian Starkey 2013-2014 <stark3y@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of version 2 of the GNU General Public License as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package basic2d

import (
    "fmt"
    "image"
    "image/draw"
    "image/color"
)

const debug bool = false

func Box(dst draw.Image, r image.Rectangle, w int, c color.Color) {
    mask := image.NewAlpha(dst.Bounds())
    black := color.Alpha{0}

    // Opaque mask
    draw.Draw(mask, mask.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)
    // Transparent rectangle full size of box
    draw.Draw(mask, r, &image.Uniform{color.Alpha{255}}, image.ZP, draw.Src)
    // Opaque rectangle inside box (leaving only border transparent)
    ir := image.Rectangle{r.Min.Add(image.Point{w, w}), r.Max.Sub(image.Point{w, w})}
    if (debug) {
        fmt.Printf("Outer Rect: %v, Inner Rect %v\n", r, ir)
    }
    draw.Draw(mask, ir, &image.Uniform{black}, image.ZP, draw.Src)

    draw.DrawMask(dst, dst.Bounds(), &image.Uniform{c}, image.ZP, mask, image.ZP, draw.Over)
}

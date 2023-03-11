package tulip

const (
	// Block: General Punctation U+2000-U+206F (http://unicode.org/charts/PDF/U2000.pdf)
	SemigraphicsHorizontalEllipsis rune = '\u2026' // …

	// Block: Box Drawing U+2500-U+257F (http://unicode.org/charts/PDF/U2500.pdf)
	BoxDrawingsLightHorizontal                    rune = '\u2500' // ─
	BoxDrawingsHeavyHorizontal                    rune = '\u2501' // ━
	BoxDrawingsLightVertical                      rune = '\u2502' // │
	BoxDrawingsHeavyVertical                      rune = '\u2503' // ┃
	BoxDrawingsLightTripleDashHorizontal          rune = '\u2504' // ┄
	BoxDrawingsHeavyTripleDashHorizontal          rune = '\u2505' // ┅
	BoxDrawingsLightTripleDashVertical            rune = '\u2506' // ┆
	BoxDrawingsHeavyTripleDashVertical            rune = '\u2507' // ┇
	BoxDrawingsLightQuadrupleDashHorizontal       rune = '\u2508' // ┈
	BoxDrawingsHeavyQuadrupleDashHorizontal       rune = '\u2509' // ┉
	BoxDrawingsLightQuadrupleDashVertical         rune = '\u250a' // ┊
	BoxDrawingsHeavyQuadrupleDashVertical         rune = '\u250b' // ┋
	BoxDrawingsLightDownAndRight                  rune = '\u250c' // ┌
	BoxDrawingsDownLighAndRightHeavy              rune = '\u250d' // ┍
	BoxDrawingsDownHeavyAndRightLight             rune = '\u250e' // ┎
	BoxDrawingsHeavyDownAndRight                  rune = '\u250f' // ┏
	BoxDrawingsLightDownAndLeft                   rune = '\u2510' // ┐
	BoxDrawingsDownLighAndLeftHeavy               rune = '\u2511' // ┑
	BoxDrawingsDownHeavyAndLeftLight              rune = '\u2512' // ┒
	BoxDrawingsHeavyDownAndLeft                   rune = '\u2513' // ┓
	BoxDrawingsLightUpAndRight                    rune = '\u2514' // └
	BoxDrawingsUpLightAndRightHeavy               rune = '\u2515' // ┕
	BoxDrawingsUpHeavyAndRightLight               rune = '\u2516' // ┖
	BoxDrawingsHeavyUpAndRight                    rune = '\u2517' // ┗
	BoxDrawingsLightUpAndLeft                     rune = '\u2518' // ┘
	BoxDrawingsUpLightAndLeftHeavy                rune = '\u2519' // ┙
	BoxDrawingsUpHeavyAndLeftLight                rune = '\u251a' // ┚
	BoxDrawingsHeavyUpAndLeft                     rune = '\u251b' // ┛
	BoxDrawingsLightVerticalAndRight              rune = '\u251c' // ├
	BoxDrawingsVerticalLightAndRightHeavy         rune = '\u251d' // ┝
	BoxDrawingsUpHeavyAndRightDownLight           rune = '\u251e' // ┞
	BoxDrawingsDownHeacyAndRightUpLight           rune = '\u251f' // ┟
	BoxDrawingsVerticalHeavyAndRightLight         rune = '\u2520' // ┠
	BoxDrawingsDownLightAnbdRightUpHeavy          rune = '\u2521' // ┡
	BoxDrawingsUpLightAndRightDownHeavy           rune = '\u2522' // ┢
	BoxDrawingsHeavyVerticalAndRight              rune = '\u2523' // ┣
	BoxDrawingsLightVerticalAndLeft               rune = '\u2524' // ┤
	BoxDrawingsVerticalLightAndLeftHeavy          rune = '\u2525' // ┥
	BoxDrawingsUpHeavyAndLeftDownLight            rune = '\u2526' // ┦
	BoxDrawingsDownHeavyAndLeftUpLight            rune = '\u2527' // ┧
	BoxDrawingsVerticalheavyAndLeftLight          rune = '\u2528' // ┨
	BoxDrawingsDownLightAndLeftUpHeavy            rune = '\u2529' // ┨
	BoxDrawingsUpLightAndLeftDownHeavy            rune = '\u252a' // ┪
	BoxDrawingsHeavyVerticalAndLeft               rune = '\u252b' // ┫
	BoxDrawingsLightDownAndHorizontal             rune = '\u252c' // ┬
	BoxDrawingsLeftHeavyAndRightDownLight         rune = '\u252d' // ┭
	BoxDrawingsRightHeavyAndLeftDownLight         rune = '\u252e' // ┮
	BoxDrawingsDownLightAndHorizontalHeavy        rune = '\u252f' // ┯
	BoxDrawingsDownHeavyAndHorizontalLight        rune = '\u2530' // ┰
	BoxDrawingsRightLightAndLeftDownHeavy         rune = '\u2531' // ┱
	BoxDrawingsLeftLightAndRightDownHeavy         rune = '\u2532' // ┲
	BoxDrawingsHeavyDownAndHorizontal             rune = '\u2533' // ┳
	BoxDrawingsLightUpAndHorizontal               rune = '\u2534' // ┴
	BoxDrawingsLeftHeavyAndRightUpLight           rune = '\u2535' // ┵
	BoxDrawingsRightHeavyAndLeftUpLight           rune = '\u2536' // ┶
	BoxDrawingsUpLightAndHorizontalHeavy          rune = '\u2537' // ┷
	BoxDrawingsUpHeavyAndHorizontalLight          rune = '\u2538' // ┸
	BoxDrawingsRightLightAndLeftUpHeavy           rune = '\u2539' // ┹
	BoxDrawingsLeftLightAndRightUpHeavy           rune = '\u253a' // ┺
	BoxDrawingsHeavyUpAndHorizontal               rune = '\u253b' // ┻
	BoxDrawingsLightVerticalAndHorizontal         rune = '\u253c' // ┼
	BoxDrawingsLeftHeavyAndRightVerticalLight     rune = '\u253d' // ┽
	BoxDrawingsRightHeavyAndLeftVerticalLight     rune = '\u253e' // ┾
	BoxDrawingsVerticalLightAndHorizontalHeavy    rune = '\u253f' // ┿
	BoxDrawingsUpHeavyAndDownHorizontalLight      rune = '\u2540' // ╀
	BoxDrawingsDownHeavyAndUpHorizontalLight      rune = '\u2541' // ╁
	BoxDrawingsVerticalHeavyAndHorizontalLight    rune = '\u2542' // ╂
	BoxDrawingsLeftUpHeavyAndRightDownLight       rune = '\u2543' // ╃
	BoxDrawingsRightUpHeavyAndLeftDownLight       rune = '\u2544' // ╄
	BoxDrawingsLeftDownHeavyAndRightUpLight       rune = '\u2545' // ╅
	BoxDrawingsRightDownHeavyAndLeftUpLight       rune = '\u2546' // ╆
	BoxDrawingsDownLightAndUpHorizontalHeavy      rune = '\u2547' // ╇
	BoxDrawingsUpLightAndDownHorizontalHeavy      rune = '\u2548' // ╈
	BoxDrawingsRightLightAndLeftVerticalHeavy     rune = '\u2549' // ╉
	BoxDrawingsLeftLightAndRightVerticalHeavy     rune = '\u254a' // ╊
	BoxDrawingsHeavyVerticalAndHorizontal         rune = '\u254b' // ╋
	BoxDrawingsLightDoubleDashHorizontal          rune = '\u254c' // ╌
	BoxDrawingsHeavyDoubleDashHorizontal          rune = '\u254d' // ╍
	BoxDrawingsLightDoubleDashVertical            rune = '\u254e' // ╎
	BoxDrawingsHeavyDoubleDashVertical            rune = '\u254f' // ╏
	BoxDrawingsDoubleHorizontal                   rune = '\u2550' // ═
	BoxDrawingsDoubleVertical                     rune = '\u2551' // ║
	BoxDrawingsDownSingleAndRightDouble           rune = '\u2552' // ╒
	BoxDrawingsDownDoubleAndRightSingle           rune = '\u2553' // ╓
	BoxDrawingsDoubleDownAndRight                 rune = '\u2554' // ╔
	BoxDrawingsDownSingleAndLeftDouble            rune = '\u2555' // ╕
	BoxDrawingsDownDoubleAndLeftSingle            rune = '\u2556' // ╖
	BoxDrawingsDoubleDownAndLeft                  rune = '\u2557' // ╗
	BoxDrawingsUpSingleAndRightDouble             rune = '\u2558' // ╘
	BoxDrawingsUpDoubleAndRightSingle             rune = '\u2559' // ╙
	BoxDrawingsDoubleUpAndRight                   rune = '\u255a' // ╚
	BoxDrawingsUpSingleAndLeftDouble              rune = '\u255b' // ╛
	BoxDrawingsUpDobuleAndLeftSingle              rune = '\u255c' // ╜
	BoxDrawingsDoubleUpAndLeft                    rune = '\u255d' // ╝
	BoxDrawingsVerticalSingleAndRightDouble       rune = '\u255e' // ╞
	BoxDrawingsVerticalDoubleAndRightSingle       rune = '\u255f' // ╟
	BoxDrawingsDoubleVerticalAndRight             rune = '\u2560' // ╠
	BoxDrawingsVerticalSingleAndLeftDouble        rune = '\u2561' // ╡
	BoxDrawingsVerticalDoubleAndLeftSingle        rune = '\u2562' // ╢
	BoxDrawingsDoubleVerticalAndLeft              rune = '\u2563' // ╣
	BoxDrawingsDownSingleAndHorizontalDouble      rune = '\u2564' // ╤
	BoxDrawingsDownDoubleAndHorizontalSingle      rune = '\u2565' // ╥
	BoxDrawingsDoubleDownAndHorizontal            rune = '\u2566' // ╦
	BoxDrawingsUpSingleAndHorizontalDouble        rune = '\u2567' // ╧
	BoxDrawingsUpDoubleAndHorizontalSingle        rune = '\u2568' // ╨
	BoxDrawingsDoubleUpAndHorizontal              rune = '\u2569' // ╩
	BoxDrawingsVerticalSingleAndHorizontalDouble  rune = '\u256a' // ╪
	BoxDrawingsVerticalDoubleAndHorizontalSingle  rune = '\u256b' // ╫
	BoxDrawingsDoubleVerticalAndHorizontal        rune = '\u256c' // ╬
	BoxDrawingsLightArcDownAndRight               rune = '\u256d' // ╭
	BoxDrawingsLightArcDownAndLeft                rune = '\u256e' // ╮
	BoxDrawingsLightArcUpAndLeft                  rune = '\u256f' // ╯
	BoxDrawingsLightArcUpAndRight                 rune = '\u2570' // ╰
	BoxDrawingsLightDiagonalUpperRightToLowerLeft rune = '\u2571' // ╱
	BoxDrawingsLightDiagonalUpperLeftToLowerRight rune = '\u2572' // ╲
	BoxDrawingsLightDiagonalCross                 rune = '\u2573' // ╳
	BoxDrawingsLightLeft                          rune = '\u2574' // ╴
	BoxDrawingsLightUp                            rune = '\u2575' // ╵
	BoxDrawingsLightRight                         rune = '\u2576' // ╶
	BoxDrawingsLightDown                          rune = '\u2577' // ╷
	BoxDrawingsHeavyLeft                          rune = '\u2578' // ╸
	BoxDrawingsHeavyUp                            rune = '\u2579' // ╹
	BoxDrawingsHeavyRight                         rune = '\u257a' // ╺
	BoxDrawingsHeavyDown                          rune = '\u257b' // ╻
	BoxDrawingsLightLeftAndHeavyRight             rune = '\u257c' // ╼
	BoxDrawingsLightUpAndHeavyDown                rune = '\u257d' // ╽
	BoxDrawingsHeavyLeftAndLightRight             rune = '\u257e' // ╾
	BoxDrawingsHeavyUpAndLightDown                rune = '\u257f' // ╿
)

// Block Elements: https://www.unicode.org/charts/PDF/U2580.pdf
const (
	BlockUpperHalf                                   rune = '\u2580' // ▀ UPPER HALF BLOCK
	BlockLowerOneEighth                              rune = '\u2581' // ▁ LOWER ONE EIGHTH BLOCK
	BlockLowerOneQuarter                             rune = '\u2582' // ▂ LOWER ONE QUARTER BLOCK
	BlockLowerThreeEighths                           rune = '\u2583' // ▃ LOWER THREE EIGHTHS BLOCK
	BlockLowerHalf                                   rune = '\u2584' // ▄ LOWER HALF BLOCK
	BlockLowerFiveEighths                            rune = '\u2585' // ▅ LOWER FIVE EIGHTHS BLOCK
	BlockLowerThreeQuarters                          rune = '\u2586' // ▆ LOWER THREE QUARTERS BLOCK
	BlockLowerSevenEighths                           rune = '\u2587' // ▇ LOWER SEVEN EIGHTHS BLOCK
	BlockFull                                        rune = '\u2588' // █ FULL BLOCK = solid → 25A0 ■  black square
	BlockLeftSevenEighths                            rune = '\u2589' // ▉ LEFT SEVEN EIGHTHS BLOCK
	BlockLeftThreeQuarters                           rune = '\u258A' // ▊ LEFT THREE QUARTERS BLOCK
	BlockLeftFiveEighths                             rune = '\u258B' // ▋ LEFT FIVE EIGHTHS BLOCK
	BlockLeftHalf                                    rune = '\u258C' // ▌ LEFT HALF BLOCK
	BlockLeftThreeEights                             rune = '\u258D' // ▍ LEFT THREE EIGHTHS BLOCK
	BlockLeftOneQuarter                              rune = '\u258E' // ▎ LEFT ONE QUARTER BLOCK
	BlockLeftOneEghth                                rune = '\u258F' // ▏ LEFT ONE EIGHTH BLOCK
	BlockRightHalf                                   rune = '\u2590' // ▐ RIGHT HALF BLOCK
	ShadeLight                                       rune = '\u2591' // ░ LIGHT SHADE • 25%
	ShadeMedium                                      rune = '\u2592' // ▒ MEDIUM SHADE = speckles fill, dotted fill • 50% • used in mapping to cp949 → 1FB90 🮐  inverse medium shade
	DarkShade                                        rune = '\u2593' // ▓ DARK SHADE • 75%
	BlockUpperOneEighth                              rune = '\u2594' // ▔ UPPER ONE EIGHTH BLOCK
	BlockRightOneEighth                              rune = '\u2595' // ▕ RIGHT ONE EIGHTH BLOCK
	BlockQuadrantLowerLeft                           rune = '\u2596' // ▖ QUADRANT LOWER LEFT
	BlockQuadrantLowerRight                          rune = '\u2597' // ▗ QUADRANT LOWER RIGHT
	BlockQuadrantUpperLeft                           rune = '\u2598' // ▘ QUADRANT UPPER LEFT
	BlockQuadrantUpperLeftAndLowerLeftAndLowerRight  rune = '\u2599' // ▙ QUADRANT UPPER LEFT AND LOWER LEFT AND LOWER RIGHT
	BlockQuadrantUpperLeftAndLowerRight              rune = '\u259A' // ▚ QUADRANT UPPER LEFT AND LOWER RIGHT → 1F67F 🙿  reverse checker board → 1FB95 🮕  checker board fill
	BlockQuadrantUpperLeftAndUpperRightAndLowerLeft  rune = '\u259B' // ▛ QUADRANT UPPER LEFT AND UPPER RIGHT AND LOWER LEFT
	BlockQuadrantUpperLeftAndUpperRightAndLowerRight rune = '\u259C' // ▜ QUADRANT UPPER LEFT AND UPPER RIGHT AND LOWER RIGHT
	BlockQuadrantUpperRight                          rune = '\u259D' // ▝ QUADRANT UPPER RIGHT
	BlockQuadrantUpperRightAndLowerLeft              rune = '\u259E' // ▞ QUADRANT UPPER RIGHT AND LOWER LEFT → 1F67E 🙾  checker board → 1FB96 🮖  inverse checker board fill
	BlockQuadrantUpperAndLowerLeftAndLowerRight      rune = '\u259F' // ▟ QUADRANT UPPER RIGHT AND LOWER LEFT AND LOWER RIGHT
)

// From Geometric Shapes https://www.unicode.org/charts/PDF/U25A0.pdf
const (
	WhiteCircle                   rune = '\u25CB' // ○ WHITE CIRCLE
	BlackCircle                   rune = '\u25CF' // ● BLACK CIRCLE
	UpperLeftQuadrantCircularArc  rune = '\u25DC' // ◜ UPPER LEFT QUADRANT CIRCULAR ARC
	UpperRightQuadrantCircularArc rune = '\u25DD' // ◝ UPPER RIGHT QUADRANT CIRCULAR ARC
	LowerRightQuadrantCircularArc rune = '\u25DE' // ◞ LOWER RIGHT QUADRANT CIRCULAR ARC
	LowerLeftQuadrantCircularArc  rune = '\u25DF' // ◟ LOWER LEFT QUADRANT CIRCULAR ARC
	UpperHalfCircle               rune = '\u25E0' // ◠ UPPER HALF CIRCLE → 2312 ⌒  arc
	LowerHalfCircle               rune = '\u25E1' // ◡ LOWER HALF CIRCLE
	BlackLowerRightTriangle       rune = '\u25E2' // ◢ BLACK LOWER RIGHT TRIANGLE → 1FB9E 🮞  lower right triangular medium shade
	BlackLowerLeftTriangle        rune = '\u25E3' // ◣ BLACK LOWER LEFT TRIANGLE  → 1FB9F 🮟  lower left triangular medium shade
	BlackUpperLeftTriangle        rune = '\u25E4' // ◤ BLACK UPPER LEFT TRIANGLE  → 1FB9C 🮜  upper left triangular medium shade
	BlackUpperRighttTriangle      rune = '\u25E5' // ◥ BLACK UPPER RIGHT TRIANGLE → 1FB9D �
	BlackUpPointingTriangle       rune = '\u25B2' // ▲ BLACK UP-POINTING TRIANGLE
	BlackRightPointingTriangle    rune = '\u25B6' // ▶ BLACK RIGHT-POINTING TRIANGLE
	BlackDownPointingTriangle     rune = '\u25BC' // ▼ BLACK DOWN-POINTING TRIANGLE
	BlackLeftPointingTriangle     rune = '\u25C0' // ◀ BLACK LEFT-POINTING TRIANGLE
)

// From Supplemental Arrows-C https://www.unicode.org/charts/PDF/U1F800.pdf
const (
	LefwardsHeavyArrow   rune = '\U0001F844' // 🡄 LEFTWARDS HEAVY ARROW
	UpwardsHeavyArrow    rune = '\U0001F845' // 🡅 UPWARDS HEAVY ARROW
	RightwardsHeavyArrow rune = '\U0001F846' // 🡆 RIGHTWARDS HEAVY ARROW
	DownwardsHeavyArrow  rune = '\U0001F847' // 🡇 DOWNWARDS HEAVY ARROW
)

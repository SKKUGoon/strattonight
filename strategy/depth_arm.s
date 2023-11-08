// +build arm,!noasm

// func addVectors(dest, src1, src2 []float32)
TEXT ·addVectors(SB), NOSPLIT, $0-28
    // Assuming len(dest) == len(src1) == len(src2) and all are multiples of 4
    MOVW    dest_len+4(FP), R0  // R0 = len(dest)
    MOVW    dest_base+0(FP), R1 // R1 = &dest[0]
    MOVW    src1_base+8(FP), R2 // R2 = &src1[0]
    MOVW    src2_base+12(FP), R3 // R3 = &src2[0]

    // R4 = loop counter initialized to 0
    MOVW    $0, R4

    // Loop start (for i := 0; i < len(dest); i += 4)
LOOP:
    // Using NEON instructions to add vectors
    VLD1.F32 {Q0}, [R2]!  // Load 4 float32s from src1 into Q0, increment R2
    VLD1.F32 {Q1}, [R3]!  // Load 4 float32s from src2 into Q1, increment R3
    VADD.F32 Q0, Q0, Q1   // Add the vectors in Q0 and Q1, store result in Q0
    VST1.F32 {Q0}, [R1]!  // Store the result from Q0 into dest, increment R1

    ADDS    R4, R4, #4     // Increment loop counter by 4
    CMP     R4, R0         // Compare loop counter with len(dest)
    BLT     LOOP           // If R4 < R0, loop

    RET

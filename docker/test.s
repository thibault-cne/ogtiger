main:
	/* Display allocation */
	MOV r0, #4
	BL malloc
	MOV r10, r0

	BL blk_4_1
	B _exit

blk_4_1:
	/* Display edit on entering a block */
	LDR r0, [r10, #-4]
	STMFD r13!, {r11, r14, r0}
	MOV r11, r13
	STR r11, [r10, #-4]

	/* Declarations de x */
	MOV r8, #1
	MOV r0, #0
	SUB r8, r0, r8
	STMFD r13!, {r8}

	/* Call the print fn */
	/* Load the format string parameter for the print function */
	LDR r0, =format_debug_int
	STMFD r13!, {r0}

	/* Use the static chain to get the value of x */
	LDR r0, [r10, #-4]
	ADD r0, r0, #-4
	LDR r8, [r0, #0]

	/* Arg x of function print */
	STMFD r13!, {r8}

	BL _print

	/* Remove args */
	ADD r13, r13, #8

	/* Remove declaration from stack */
	ADD r13, r13, #4

	/* Display edit on exiting a block */
	LDMFD r13!, {r0}
	STR r0, [r10, #-4]
	LDMFD r13!, {r15, r11}


_print:
	STMFD r13!, {r11, r14}
	MOV r11, r13
	STMFD r13!, {r0, r1}

	LDR r0, [r11, #12]
	LDR r1, [r11, #8]

	BL printf

	MOV r0, #0

	BL fflush

	LDMFD r13!, {r0, r1}
	LDMFD r13!, {r15, r11}
	
_exit:
	MOV r0, #0
	MOV r7, #1
	SWI 0

.data
	format_str:			.ascii		"%s\0"
	format_int:			.ascii		"%d\0"
	format_debug_int:	.ascii		"debug: %d\n\0"
	format_debug_x:		.ascii		"debug: %x\n\0"
	format_debug_addr:	.ascii		"debug: %p\n\0"
	scanf_int:			.ascii		"%d"

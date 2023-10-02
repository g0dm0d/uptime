import type { CustomThemeConfig } from '@skeletonlabs/tw-plugin';

export const mainTheme: CustomThemeConfig = {
    name: 'site-main-theme',
    properties: {
		// =~= Theme Properties =~=
		"--theme-font-family-base": `system-ui`,
		"--theme-font-family-heading": `system-ui`,
		"--theme-font-color-base": "var(--color-surface-800)",
		"--theme-font-color-dark": "var(--color-surface-100)",
		"--theme-rounded-base": "9999px",
		"--theme-rounded-container": "8px",
		"--theme-border-base": "2px",
		// =~= Theme On-X Colors =~=
		"--on-primary": "var(--color-surface-500)",
		"--on-secondary": "var(--color-surface-900)",
		"--on-tertiary": "var(--color-surface-900)",
		"--on-success": "var(--color-surface-900)",
		"--on-warning": "var(--color-surface-900)",
		"--on-error": "var(--color-surface-900)",
		"--on-surface": "var(--color-surface-50)",
		// =~= Theme Colors  =~=
		// primary | #a6d189 
		"--color-primary-50": "242 248 237", // #f2f8ed
		"--color-primary-100": "237 246 231", // #edf6e7
		"--color-primary-200": "233 244 226", // #e9f4e2
		"--color-primary-300": "219 237 208", // #dbedd0
		"--color-primary-400": "193 223 172", // #c1dfac
		"--color-primary-500": "166 209 137", // #a6d189
		"--color-primary-600": "149 188 123", // #95bc7b
		"--color-primary-700": "125 157 103", // #7d9d67
		"--color-primary-800": "100 125 82", // #647d52
		"--color-primary-900": "81 102 67", // #516643
		// secondary | #99d1db 
		"--color-secondary-50": "240 248 250", // #f0f8fa
		"--color-secondary-100": "235 246 248", // #ebf6f8
		"--color-secondary-200": "230 244 246", // #e6f4f6
		"--color-secondary-300": "214 237 241", // #d6edf1
		"--color-secondary-400": "184 223 230", // #b8dfe6
		"--color-secondary-500": "153 209 219", // #99d1db
		"--color-secondary-600": "138 188 197", // #8abcc5
		"--color-secondary-700": "115 157 164", // #739da4
		"--color-secondary-800": "92 125 131", // #5c7d83
		"--color-secondary-900": "75 102 107", // #4b666b
		// tertiary | #81c8be 
		"--color-tertiary-50": "236 247 245", // #ecf7f5
		"--color-tertiary-100": "230 244 242", // #e6f4f2
		"--color-tertiary-200": "224 241 239", // #e0f1ef
		"--color-tertiary-300": "205 233 229", // #cde9e5
		"--color-tertiary-400": "167 217 210", // #a7d9d2
		"--color-tertiary-500": "129 200 190", // #81c8be
		"--color-tertiary-600": "116 180 171", // #74b4ab
		"--color-tertiary-700": "97 150 143", // #61968f
		"--color-tertiary-800": "77 120 114", // #4d7872
		"--color-tertiary-900": "63 98 93", // #3f625d
		// success | #84cc16 
		"--color-success-50": "237 247 220", // #edf7dc
		"--color-success-100": "230 245 208", // #e6f5d0
		"--color-success-200": "224 242 197", // #e0f2c5
		"--color-success-300": "206 235 162", // #ceeba2
		"--color-success-400": "169 219 92", // #a9db5c
		"--color-success-500": "132 204 22", // #84cc16
		"--color-success-600": "119 184 20", // #77b814
		"--color-success-700": "99 153 17", // #639911
		"--color-success-800": "79 122 13", // #4f7a0d
		"--color-success-900": "65 100 11", // #41640b
		// warning | #e5a50a 
		"--color-warning-50": "251 242 218", // #fbf2da
		"--color-warning-100": "250 237 206", // #faedce
		"--color-warning-200": "249 233 194", // #f9e9c2
		"--color-warning-300": "245 219 157", // #f5db9d
		"--color-warning-400": "237 192 84", // #edc054
		"--color-warning-500": "229 165 10", // #e5a50a
		"--color-warning-600": "206 149 9", // #ce9509
		"--color-warning-700": "172 124 8", // #ac7c08
		"--color-warning-800": "137 99 6", // #896306
		"--color-warning-900": "112 81 5", // #705105
		// error | #fa383e 
		"--color-error-50": "254 225 226", // #fee1e2
		"--color-error-100": "254 215 216", // #fed7d8
		"--color-error-200": "254 205 207", // #fecdcf
		"--color-error-300": "253 175 178", // #fdafb2
		"--color-error-400": "252 116 120", // #fc7478
		"--color-error-500": "250 56 62", // #fa383e
		"--color-error-600": "225 50 56", // #e13238
		"--color-error-700": "188 42 47", // #bc2a2f
		"--color-error-800": "150 34 37", // #962225
		"--color-error-900": "123 27 30", // #7b1b1e
		// surface | #292524 
		"--color-surface-50": "223 222 222", // #dfdede
		"--color-surface-100": "212 211 211", // #d4d3d3
		"--color-surface-200": "202 201 200", // #cac9c8
		"--color-surface-300": "169 168 167", // #a9a8a7
		"--color-surface-400": "105 102 102", // #696666
		"--color-surface-500": "41 37 36", // #292524
		"--color-surface-600": "37 33 32", // #252120
		"--color-surface-700": "31 28 27", // #1f1c1b
		"--color-surface-800": "25 22 22", // #191616
		"--color-surface-900": "20 18 18", // #141212
		
	}
}
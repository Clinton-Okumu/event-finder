/**
 * Below are the colors that are used in the app. The colors are defined in the light and dark mode.
 * Colors match frontend's OKLCH theme exactly
 */

import { Platform } from 'react-native';

export const Colors = {
  light: {
    text: '#030712',
    background: '#ffffff',
    tint: '#9336ea',
    icon: '#6a7282',
    tabIconDefault: '#9ca3af',
    tabIconSelected: '#9336ea',
    card: '#ffffff',
    border: '#e5e7eb',
    input: '#e5e7eb',
    primary: '#9336ea',
    primaryForeground: '#faf7fe',
    secondary: '#f4f4f5',
    secondaryForeground: '#18181b',
    muted: '#f3f4f6',
    mutedForeground: '#6a7282',
    accent: '#f3f4f6',
    accentForeground: '#101828',
    destructive: '#e7000b',
    destructiveForeground: '#f8f8f8',
    ring: '#99a1af',
  },
  dark: {
    text: '#f9fafb',
    background: '#030712',
    tint: '#a957f7',
    icon: '#99a1af',
    tabIconDefault: '#6a7282',
    tabIconSelected: '#a957f7',
    card: '#101828',
    border: '#ffffff',
    input: '#ffffff',
    primary: '#a957f7',
    primaryForeground: '#faf7fe',
    secondary: '#27272a',
    secondaryForeground: '#fafafa',
    muted: '#1e2939',
    mutedForeground: '#99a1af',
    accent: '#1e2939',
    accentForeground: '#f9fafb',
    destructive: '#ff6467',
    destructiveForeground: '#f8f8f8',
    ring: '#6a7282',
  },
};

export const Fonts = Platform.select({
  ios: {
    /** iOS `UIFontDescriptorSystemDesignDefault` */
    sans: 'system-ui',
    /** iOS `UIFontDescriptorSystemDesignSerif` */
    serif: 'ui-serif',
    /** iOS `UIFontDescriptorSystemDesignRounded` */
    rounded: 'ui-rounded',
    /** iOS `UIFontDescriptorSystemDesignMonospaced` */
    mono: 'ui-monospace',
  },
  default: {
    sans: 'normal',
    serif: 'serif',
    rounded: 'normal',
    mono: 'monospace',
  },
  web: {
    sans: "system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif",
    serif: "Georgia, 'Times New Roman', serif",
    rounded: "'SF Pro Rounded', 'Hiragino Maru Gothic ProN', Meiryo, 'MS PGothic', sans-serif",
    mono: "SFMono-Regular, Menlo, Monaco, Consolas, 'Liberation Mono', 'Courier New', monospace",
  },
});

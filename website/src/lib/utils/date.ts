export function formatDate(dateString: string): string {
    const date = new Date(dateString);
    const options: Intl.DateTimeFormatOptions = { hour:'numeric', minute:'2-digit', formatMatcher:'basic' };
    return date.toLocaleString('en-US', options);
}
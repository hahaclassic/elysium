# Types of keys

Key is a string like 'KEY-' + random uuid (in upper case)

Example
```
KEY-C9E11A7C-23FD-4E29-9D9B-23D75C3C5DF8
```

## Access levels
1. **Owner** (All possible actions with the folder and its contents are available)
2. **Editor** (Add/delete links (always after confirmation))
3. **ConfirmedReader** (Reading only after confirmation)
4. **Reader** (Reading only)
5. **Suspected** (Last chance to gain access. In case of another refusal, the status will change to banned)
6. **Banned** (The folder owner does not receive requests from banned users. Ofc user doesn't have any access to this folder)

You can create keys only for Readers, Verified Readers, and Editors.


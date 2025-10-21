### Starting Local SMTP Server for Testing
``` zsh
docker run -d \
--restart unless-stopped \
--name=mailpit \
-p 8025:8025 \
-p 1025:1025 \
axllent/mailpit
```

### Future Improvements
- Store the contact lists into DB
- Add segments
- Multiple Campaign
- Handle email fails
- Customize Email Template
- Schedule Campaign
- Build a Frontend
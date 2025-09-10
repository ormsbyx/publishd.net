# Deployment Guide - Render

This guide walks through deploying Publishd to Render's free tier.

## Prerequisites

1. **Render Account**: Sign up at https://render.com
2. **GitHub Repository**: Push your code to GitHub
3. **Domain** (optional): Point publishd.net to your Render service

## Step 1: Create Render Account

1. Go to https://render.com
2. Click "Get Started for Free"
3. Sign up with GitHub (recommended)
4. Connect your GitHub account

## Step 2: Deploy Database

1. In Render dashboard, click "New +"
2. Select "PostgreSQL"
3. Configure:
   - **Name**: `publishd-db`
   - **Database**: `publishd`
   - **User**: `publishd_user`
   - **Plan**: Free
4. Click "Create Database"
5. Wait for database to deploy (~2-3 minutes)

## Step 3: Deploy Web Service

1. In Render dashboard, click "New +"
2. Select "Web Service"
3. Connect your GitHub repository (`published.net`)
4. Configure:
   - **Name**: `publishd-web`
   - **Runtime**: Go
   - **Build Command**: `go build -o bin/publishd .`
   - **Start Command**: `./bin/publishd`
   - **Plan**: Free

## Step 4: Environment Variables

The `render.yaml` file automatically configures these, but you can also set them manually:

```
DB_HOST=<from database service>
DB_PORT=<from database service>  
DB_NAME=publishd
DB_USER=<from database service>
DB_PASSWORD=<from database service>
DB_SSLMODE=require
GIN_MODE=release
```

## Step 5: Deploy

1. Click "Create Web Service"
2. Render will:
   - Clone your repository
   - Build the Go application
   - Start the server
   - Run database migrations automatically

## Step 6: Access Your Site

1. Render provides a URL like: `https://publishd-web.onrender.com`
2. Visit `/health` to verify deployment
3. Visit `/admin` to create your first story
4. Visit `/` to see your published stories

## Step 7: Custom Domain (Optional)

1. In Render dashboard, go to your web service
2. Click "Settings" → "Custom Domains"
3. Add `publishd.net`
4. Update DNS records:
   ```
   Type: CNAME
   Name: @
   Value: publishd-web.onrender.com
   ```

## Troubleshooting

### Build Failures
- Check Go version compatibility
- Ensure all dependencies are in `go.mod`
- Check build logs in Render dashboard

### Database Connection Issues
- Verify environment variables are set
- Check database service is running
- Ensure SSL mode is set to `require`

### Template/Static File Issues
- Verify `web/` directory is copied correctly
- Check file paths in Dockerfile
- Ensure templates load in production mode

## Free Tier Limitations

- **Web Service**: 750 hours/month, sleeps after 15 min of inactivity
- **Database**: 1GB storage, shared CPU
- **Bandwidth**: 100GB/month

## Monitoring

- **Logs**: Available in Render dashboard
- **Health Check**: `/health` endpoint monitors database connection
- **Uptime**: Render provides basic monitoring

## Production Considerations

For production use, consider:
- Upgrading to paid plans for better performance
- Adding Redis for session management
- Implementing proper error logging
- Setting up monitoring and alerting
- Adding backup strategies for database